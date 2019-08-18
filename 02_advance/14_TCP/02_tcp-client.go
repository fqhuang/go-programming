package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
从键盘获取数据发给server
*/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001") //server IP
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	go handleStdinData(conn)

	handleServerData(conn)

}

func handleStdinData(conn net.Conn) {
	buff := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buff) /////

		if err != nil && err != io.EOF {
			fmt.Println("os.Stdin.Read err:", err)
			return
		}

		fmt.Println("Stdin.Read:", string(buff[:n]))
		conn.Write(buff[:n])
	}
}

func handleServerData(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 4096)
	for {
		n, err := conn.Read(buff)
		if n == 0 {
			fmt.Println("server already closed!")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("os.Stdin.Read err:", err)
			return
		}

		fmt.Println("recv from server ", conn.RemoteAddr().String(), string(buff[0:n]))
	}
}
