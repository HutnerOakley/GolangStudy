package servers

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// server类型
type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server对象
func NewServer(ip string, port int) *Server {
	// &Server{} = 创建结构体实例 + 拿到它的内存地址（指针）
	// Server{} = 创建结构体实例
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 启动server服务
func (s *Server) Start() {
	// 监听连接
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	// 关闭listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go s.ListenMessage()

	for {
		// 接受消息
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// 处理链接业务
		go s.Handler(conn)
	}

}

func (s *Server) Handler(conn net.Conn) {
	// ...当前链接的业务
	//fmt.Println("建立链接成功")

	user := NewUser(conn, s)

	// 用户上线，将用户加到OnlineMap中
	//s.mapLock.Lock()
	//s.OnlineMap[user.Name] = user
	//s.mapLock.Unlock()
	user.Online()

	// 广播当前用户上线消息
	//s.BroadCast(user, "已上线")

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//s.BroadCast(user, "下线")
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}

			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			// 将得到的消息进行广播
			//s.BroadCast(user, msg)

			// 用户针对msg进行消息处理
			user.DoMessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
		// 当前用户是活跃的，应该重置定时器
		// 不做任何事情，为了激活select 更新下面的定时器
		case <-time.After(time.Second * 300):
			// 已经超时
			// 将当前的User强制关闭

			user.SendMsg("你被踢了")

			// 销毁用的资源
			close(user.C)

			//关闭链接
			conn.Close()

			// 退出当前Handler
			// runtime.Goexit()
			return

		}
	}
}

// 广播消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine， 一旦有消息就发送给全部在线的user
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		// 将msg发送给全部在线的user
		s.mapLock.Lock()
		for _, user := range s.OnlineMap {
			user.C <- msg
		}
		s.mapLock.Unlock()
	}
}
