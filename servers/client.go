package servers

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string // 服务端ip
	ServerPort int    // 服务端端口
	Name       string
	conn       net.Conn

	flag int // 当前client的模式
}

// 新建客户端对象
func NewClient(ServerIp string, ServerPort int) *Client {

	// 创建客户端对象
	client := &Client{
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		flag:       999,
	}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIp, ServerPort))
	if err != nil {
		fmt.Println("net dial error", err)
		return nil
	}
	client.conn = conn

	return client
}

// 获取用户输入的模式
func (c *Client) menu() bool {

	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>> 请输入合法范围内的数字")
		return false
	}
}

// 公聊模式
func (c *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	fmt.Println(">>> 请输入聊天内容，exit退出。")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送服务器
		if chatMsg != "" {
			sendMsg := chatMsg + "\n"
			if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
				fmt.Println("conn write err", err)
			}
		}

		chatMsg = ""
		fmt.Println(">>> 请输入聊天内容，exit退出。")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (c *Client) GetUsers() {
	sendMsg := "who\n"
	if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
		fmt.Println("conn write err", err)
		return
	}
}

// 私聊模式
func (c *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	c.GetUsers()
	fmt.Println(">>>> 请求输入聊天对象(用户名)，exit退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>> 请求输入内容，exit退出")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
					fmt.Println("conn write err", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>> 请求输入内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		c.GetUsers()
		fmt.Println(">>>> 请求输入聊天对象(用户名)，exit退出")
		fmt.Scanln(&remoteName)
	}
}

// 更新用户名
func (c *Client) UpdateName() bool {

	fmt.Println(">>> 请输入用户名:")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
		fmt.Println("conn write err", err)
		return false
	}
	return true
}

// 处理server回应的消息，直接显示到标准输出即可
func (c *Client) DealResponse() {
	// 一单client.conn有数据，就直接copy到stdout标准输出上，永久阻塞建议
	io.Copy(os.Stdout, c.conn)
}

// run方法 主业务
func (c *Client) Run() {
	for c.flag != 0 {
		for c.menu() != true {
		}
		// 根据不同的模式处理不同的业务
		switch c.flag {
		case 1:
			// 公聊模式
			//fmt.Println("公聊模式选择。。。")
			c.PublicChat()
			break
		case 2:
			// 私聊模式
			//fmt.Println("私聊模式选择。。。")
			c.PrivateChat()
			break
		case 3:
			// 更新用户名
			//fmt.Println("更新用户名选择。。。")
			c.UpdateName()
			break
		}
	}
}
