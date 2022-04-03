package main

import (
	"net"
	"strings"
)

//这里要做用户上线功能,所以这里定义个用户的结构体
type User struct {
	Name   string      //用户名
	Addr   string      //用户地址
	C      chan string //通道
	conn   net.Conn    //用户本身的链接
	Server *Server     //这里整理了一下,一些发送消息的内容应该是用户发送的,所以这里生成一个server的结构,用于和server交互
}

//用户上线,要在我们程序内部生成一个用户
func NewUser(conn net.Conn, server *Server) *User {
	//这里取一个唯一值
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string), //创建管道
		conn:   conn,
		Server: server,
	}

	go user.ListenMessage()

	return user

}

//这里做个监听, 监听下用户的管道(channel),一旦有消息,就直接显示在客户端里面
func (this *User) ListenMessage() {
	for {
		//还是一个死循环监听
		msg := <-this.C //这里接受消息的内容
		this.conn.Write([]byte(msg))
	}
}

//上线方法
func (this *User) Online() {
	//用户上线业务处理
	this.Server.mapLock.Lock()
	this.Server.OnlineMap[this.Name] = this
	this.Server.mapLock.Unlock()

	this.SendAllMessage("已上线")
}

//下线方法
func (this *User) OffOnline() {
	this.Server.mapLock.Lock()
	delete(this.Server.OnlineMap, this.Name)
	this.Server.mapLock.Unlock()

	//广播下线通知
	this.SendAllMessage("已下线")
}

//查看当前在线用户
func (this *User) FindOnlineUser() {
	msg := "以下为当前在线用户:\n"
	for _, val := range this.Server.OnlineMap {
		msg = msg + val.Name + "[" + val.Addr + "]\n"
	}
	this.SendThisMessage(msg)

}

//改名
func (this *User) rename(name string) {

	//检查昵称是否存在
	if _, ok := this.Server.OnlineMap[name]; ok == true {
		this.SendThisMessage("该昵称已存在,请更换昵称")
		return
	}

	this.Server.mapLock.Lock()
	delete(this.Server.OnlineMap, this.Name)
	this.Name = name
	this.Server.OnlineMap[this.Name] = this
	this.Server.mapLock.Unlock()

	this.SendThisMessage("您的昵称已改为:" + name)
}

//给自己发送消息
func (this *User) SendThisMessage(msg string) {
	this.conn.Write([]byte(msg + "\n"))
}

//私聊
func (this *User) toUserMsg(msg string) {
	//先转换格式
	msgArr := strings.Split(msg, "|")
	if msgArr[1] == "" {
		this.SendThisMessage("您输入的对方昵称不正确,请重新输入")
		return
	}

	if msgArr[2] == "" {
		this.SendThisMessage("消息不能为空")
		return
	}

	//找有没有这用户
	toUser, ok := this.Server.OnlineMap[msgArr[1]]

	//if toUser, ok := this.Server.OnlineMap[msgArr[1]]; !ok {} 这样写是不行的,会导致if执行后取不到值

	if !ok {
		this.SendThisMessage("该用户未上线,无法私聊")
		return
	}

	toUser.SendThisMessage(this.Name + "对您说:" + msgArr[2])

}

//用户广播消息方法
func (this *User) SendAllMessage(msg string) {

	if len(msg) == 3 && msg == "who" {
		this.FindOnlineUser()
		return
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//这里的命令格式是rename|张三
		this.rename(strings.Split(msg, "|")[1])
		return
	} else if len(msg) > 7 && msg[:7] == "toUser|" {
		//这里的命令格式是toUser|张三|信息
		this.toUserMsg(msg)
		return
	}

	this.Server.broadCast(this, msg)

}
