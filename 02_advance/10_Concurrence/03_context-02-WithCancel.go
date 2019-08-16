package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go watch(ctx, i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(i, "监控退出，停止了...")
			return
		default:
			fmt.Printf("goroutine %d监控中...\n", i)
			time.Sleep(1 * time.Second)
		}
	}
}
