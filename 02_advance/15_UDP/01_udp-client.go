package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	udpConn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("udp net.Dial err:", err)
		return
	}
	defer udpConn.Close()

	go func() {
		buff := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buff)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				return
			}

			udpConn.Write(buff[:n])
		}

	}()

	slice := make([]byte, 4096)
	for {
		n, err := udpConn.Read(slice)
		if err != nil {
			fmt.Println("udpConn.Read:", err)
			return
		}

		fmt.Println("recv:", string([]byte(slice[:n])))
	}

}
