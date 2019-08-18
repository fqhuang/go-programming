package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("AAA")
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("BBB")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("CCC")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go A(x, y)
	go B(y, z)
	go C(z)

	close(x)

	time.Sleep(3 * time.Second)

}
