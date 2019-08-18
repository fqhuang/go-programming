package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //创建监控器，第一个socket套接字
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept() //创建第二个套接字，通信套接字，服务器客户端成对出现
	if err != nil {
		fmt.Println("lsitener.Accept err:", err)
		return
	}
	defer conn.Close()

	slice := make([]byte, 4096)

	for {
		n, err := conn.Read(slice)
		if n == 0 {
			fmt.Println("对端", conn.RemoteAddr().String(), "关闭.")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("recv from client", conn.RemoteAddr().String(), ":", string(slice[:n])) //slice[:n] "收多少写多少"

		conn.Write([]byte("I'm server"))
	}

}
