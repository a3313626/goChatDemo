package main

import (
	"fmt"
	"net"
)

//这里要创建一个server的结构体
type Server struct {
	Ip   string
	Port int
}

//创建下服务
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

//定义链接服务
func (this *Server) HandLer(conn net.Conn) {
	//这里用到了net组件
	//当前链接业务

	//先抛出个提示
	fmt.Println("链接创建成功")

}

//这里写个启动服务的方法
func (this *Server) Start() {
	//使用net组件开启socket服务
	//fmt.Sprintf("%s:%d" , this.Ip , this.Port) 输出内容 127.0.0.1:8030
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	//判断创建链接有没有错误,错误就抛出个错误
	if err != nil {
		fmt.Println("socket服务创建失败,错误原因:", err)
	}

	//如果代码执行到了最后,一定要关闭socket服务,这里用defer关键字做处理
	defer listener.Close()

	//用for阻塞,防止执行完毕后直接结束
	for {
		//这里做了个新建链接的监听,用于监听有没有新链接过来
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("新链接接入服务器失败,错误信息:", err)
			continue //退出本次循环
		}

		//执行成功连接后的任务
		go this.HandLer(conn)
	}

}
