package main

import (
	"flag"
	"fmt"
	"net"
)

//先生成客户端结构体
type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //这个是客户端的主要功能
}

var serverIp string
var serverPort int

//解析命令行
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址")
	flag.IntVar(&serverPort, "p", 8888, "设置服务器端口地址")
}

//写入菜单规则
func (client Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")
	fmt.Println("0:退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		//没有就跳出错误
		fmt.Println(">>>>>>>>>>>>>>>您的输入有误,请重新输入<<<<<<<<<<<<<")
		return false
	}

}

//写一个菜单执行函数
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		switch client.flag {
		case 1:
			fmt.Println("公聊模式")
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("更新用户名")
			break

		}

	}
}

//生成客户端方法
func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerPort: serverPort,
		ServerIp:   serverIp,
		flag:       99, //给一个默认值
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

	flag.Parse()

	//开始连接
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>>>>> 链接服务器失败")
		return
	}

	fmt.Println("<<<<<<<<<<<<<<<< 链接服务器成功")

	//阻塞,并且这里可以写客户端业务
	client.Run()

}
