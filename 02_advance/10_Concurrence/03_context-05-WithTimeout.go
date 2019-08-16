package main

import (
	"context"
	"fmt"
	"time"
)

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//释放资源
	defer cancel()

	ch := make(chan bool)

	go func() {
		time.Sleep(time.Second * 5)
		// time.Sleep(time.Second * 1)
		ch <- true
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout!")
	case <-ch:
		fmt.Println("chan closed.")
	}

}
func main() {
	process()
}
