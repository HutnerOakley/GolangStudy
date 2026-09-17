package servers

import (
	"net"
	"strings"
)

type User struct {
	Name string      // 姓名
	Addr string      //地址
	C    chan string // 管道
	conn net.Conn    // 链接

	server *Server // 当前用户对应的server
}

// 创建一个user对象
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线功能
func (u *User) Online() {

	// 用户上线，将用户对象加到OnlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 广播用户上线
	u.server.BroadCast(u, "已上线")

}

// 用户下线功能
func (u *User) Offline() {

	// 用户下线，将用户对象从OnlineMa中删除
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播用户下线
	u.server.BroadCast(u, "下线")
}

// 用户处理消息的业务
func (u *User) DoMessage(msg string) {
	if msg == "who" {
		// 格式：who

		// 查询用户信息,发送给对应的客户端
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			u.SendMsg(onlineMsg)
		}

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式: rename|张三

		// 从msg中取出新的name
		newName := strings.Split(msg, "|")[1]
		// 查询当前用户名是否存在
		if _, ok := u.server.OnlineMap[newName]; ok {
			u.SendMsg("用户名已经被使用\n")
		} else {
			// 将旧用户名从map中删除，将newName新增到map中
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			// 将当前用user对象的用户名更新为newName
			u.Name = newName
			u.SendMsg("您已经更新用户名:" + u.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式: to|张三|消息内容

		// 从msg中提取需要发送消息的对象名字以及消息内容
		remoteName := strings.Split(msg, "|")[1]
		content := strings.Split(msg, "|")[2]
		if remoteName == "" {
			u.SendMsg("消息格式不正确，请使用\"to|张三|你好啊\"格式。\n")
			return
		}
		if content == "" {
			u.SendMsg("没有消息内容，请重发\n")
			return
		}
		if remoteUser, ok := u.server.OnlineMap[remoteName]; ok {
			remoteUser.SendMsg(u.Name + "对您说：" + content)
		} else {
			u.SendMsg("该用户不存在")
			return
		}

	} else {
		u.server.BroadCast(u, msg)
	}
}

// 给当前User对应的客户端发送消息
func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

// 监听当前user channel的方法，一旦有消息，就直接发送给对应客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
