package main

import (
	"fmt"
	"net"
)

func main() {
	//1. 创建UDP addr结构
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:", err)
		return
	}
	//2. 创建通信socket
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer udpConn.Close() //udp只需要关闭这一个套接字

	buff := make([]byte, 4096)

	for {
		//3. 读取客户端数据
		n, clitAddr, err := udpConn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("ReadFromUDP err:", err)
			return
		}
		fmt.Println("recv:", string(buff[:n]))
		//4. 发送数据到客户端
		udpConn.WriteToUDP([]byte("i am server\n"), clitAddr)
	}

}
