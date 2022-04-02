package main

import "net"

//这里要做用户上线功能,所以这里定义个用户的结构体
type User struct {
	Name string      //用户名
	Addr string      //用户地址
	C    chan string //通道
	conn net.Conn    //用户本身的链接
}

//用户上线,要在我们程序内部生成一个用户
func NewUser(conn net.Conn) *User {
	//这里取一个唯一值
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string), //创建管道
		conn: conn,
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
