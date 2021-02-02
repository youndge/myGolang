package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	IP = "127.0.0.1"
)

func main() {
	conn, err := net.Dial("tcp", IP+":8888")
	if err != nil {
		fmt.Println("Client Dial err:", err)
	}
	for {
		//1.从终端读取一行用户输入，并准备发送给服务器
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err:", err)
		}
		//检测到输入exit退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("Client exit!")
			break
		}

		//2.将line发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Write err:", err)
		}
		fmt.Printf("Send msg success! byte:%v\n", n)
	}
}
