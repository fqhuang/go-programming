package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d号完成\n", i)
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("好了，大家都干完了，放工")
}
