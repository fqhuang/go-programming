package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "go-programming"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, key, "https://github.com/meetbetter/go-programming")

	go watch(valCtx)
	time.Sleep(time.Second * 5)
	cancel()

	time.Sleep(time.Second * 2)
}

func watch(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到value：", ctx.Value(key))
			return
		default:
			fmt.Println("github.com/meetbetter")
			time.Sleep(time.Second * 1)
		}
	}

}
