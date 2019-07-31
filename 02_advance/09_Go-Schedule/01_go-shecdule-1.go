package main

/*
	GO调度器工作原理示例代码 1
	展示了Go多个Goroutine并发执行的正常现象
*/

import (
	"fmt"

	"strconv"
	"time"
)

func main() {
	var s = make([]int, 10)
	makeThumbnails(s)
	time.Sleep(10 * time.Second)
}

func makeThumbnails(lists []int) error {
	errorChan := make(chan error)

	for i := range lists {
		go func(i int) {
			err := fmt.Errorf(strconv.Itoa(i))
			errorChan <- err
		}(i)
	}

	for range lists {
		err := <-errorChan

		if err != nil {

			fmt.Println("data:", err)
		}

	}

	return nil
}
