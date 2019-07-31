package main

/*
	Go调度器之nail goroutine现象
	请注意下面程序的输出顺序

*/
import (
	"fmt"
)

func main() {

	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		fmt.Println("--->", v)
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}
