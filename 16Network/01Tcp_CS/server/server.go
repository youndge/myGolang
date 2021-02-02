package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//1.等待客户端通过conn发消息
		//2.如果客户端一直未发消息，则协程一直阻塞在这里
		buf := make([]byte, 1024)
		fmt.Printf("Waitting for msg of colient(%v)\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Client err:%v\n", err)
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}

func main() {

	//1.设置监听，"tcp"表示使用的协议是TCP
	fmt.Println("Server is listening...")
	//2."0.0.0.0:88888"表示在本地监听8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Listen err:", err)
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("Waitting for client...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		} else {
			fmt.Printf("Accept success:conn->%v,client ip->%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
