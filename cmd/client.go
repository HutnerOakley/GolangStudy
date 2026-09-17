package main

import (
	"flag"
	"fmt"

	"go_project/servers"
)

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := servers.NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>> 链接服务器败。。。")
		return
	}
	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>链接服务器成功。。。")

	// 启动客户端的业务
	client.Run()

}
