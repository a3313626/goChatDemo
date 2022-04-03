package main

import (
	"fmt"
	"net"
)

//先生成客户端结构体
type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

//生成客户端方法
func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerPort: serverPort,
		ServerIp:   serverIp,
	}

	//使用tcp链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	return client

}

func main() {
	//开始连接
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>>>>>>>> 链接服务器失败")
		return
	}

	fmt.Println("<<<<<<<<<<<<<<<< 链接服务器成功")

	//阻塞,并且这里可以写客户端业务
	select {}

}
