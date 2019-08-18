package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000") //socket
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	slice := make([]byte, 4096)
	for {
		conn.Write([]byte("hello,I'm client"))

		n, err := conn.Read(slice)
		if n == 0 {
			fmt.Println("server socket closed")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("read from ", conn.RemoteAddr().String(), ":", string(slice[:n])) //slice[:n] "收多少写多少"!!!

		<-time.After(time.Second * 2) //使用afetr要加<-
	}
}
