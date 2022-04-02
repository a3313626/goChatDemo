package main

import (
	"fmt"
	"net"
	"sync"
)

//这里要创建一个server的结构体
type Server struct {
	Ip   string
	Port int

	//在线用户
	OnlineMap map[string]*User //这个是所有用户的在线合集
	mapLock   sync.RWMutex     //这里多了个锁

	//消息广播的channel
	Message chan string
}

//创建服务
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听message广播信息
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock() //这里加个锁
		//给所有在线用户发送消息,这里用了管道传输过去
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock() //解锁
	}
}

//广播信息方法
func (this *Server) broadCast(user *User, msg string) {
	//拼接下写入信息 这里用了用户地址,用户名:消息的格式
	sendMsg := "[" + user.Addr + "]" + "~~" + user.Name + ":" + msg + "\n"
	//这里给全局的消息发送下,把发送消息,写入message管道中
	this.Message <- sendMsg
}

//定义链接服务
func (this *Server) HandLer(conn net.Conn) {
	//这里用到了net组件
	//当前链接业务

	//先抛出个提示
	//fmt.Println("链接创建成功")

	user := NewUser(conn)

	//用户上线业务处理
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//这里做一个用户上线通知
	this.broadCast(user, "已上线")

	//阻塞下
	select {}

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

	//这里要监听下有没有人发送消息,
	go this.ListenMessage()

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
