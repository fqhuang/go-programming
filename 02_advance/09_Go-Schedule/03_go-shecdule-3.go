package main

/*
	Go调度器之nail goroutine现象
	请注意下面程序的输出顺序

*/

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var s = make([]int, 10)
	err := makeThumbnails(s)
	fmt.Println(err.Error())
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
			return err
		}
		return nil
	}

	return nil
}
