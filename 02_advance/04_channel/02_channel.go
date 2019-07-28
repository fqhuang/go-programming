package main

import (
	"fmt"
	"time"
)

func send(done, quit chan bool) {
	//等待recv完成后由send关闭
	<-done
	fmt.Println("那我关闭了哦")

	//通知main goroutine
	quit <- true
}

func recv(done chan bool) {
	time.Sleep(time.Second)
	// 通知任务已完成
	fmt.Println("对面的send可以关闭了")
	done <- true
}

func main() {
	done := make(chan bool)
	quit := make(chan bool)
	go send(done, quit)
	go recv(done)
	// 等待任务完成
	<-quit
	fmt.Println("main也结束了哦")
}
