package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close() ///

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue //不能return
		}

		go handleConnect(conn)

	}

}

func handleConnect(conn net.Conn) {
	defer conn.Close()

	slice := make([]byte, 4096)
	for {
		n, err := conn.Read(slice)
		if n == 0 {
			fmt.Println("client", conn.RemoteAddr().String(), "closed")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("recv from client:", string(slice[:n])) //slice[:n]收多少发多少
		if string(slice[:n]) == "exit\n" {                  //slice[:n]收多少发多少
			fmt.Println("closing client:", conn.RemoteAddr().String())
			return //defer already
		}

		slice1 := bytes.ToUpper(slice[:n]) //slice[:n]收多少发多少
		conn.Write(slice1)
	}

}
