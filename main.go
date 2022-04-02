package main

func main() {
	server := NewServer("127.0.0.1", 8000)
	//启动下socket服务
	server.Start()
}
