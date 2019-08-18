package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	// range 一直读取直到chan关闭，否则产生阻塞死锁
	//解决方式：
	//
	//a. 显式关闭 channel；
	//
	//b. for range chan放进子协程，主协程 sleep 等待时间后退出；
	for v := range ch {
		fmt.Println(v)
	}
}
