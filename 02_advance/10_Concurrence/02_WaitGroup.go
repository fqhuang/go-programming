package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d号完成\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}
