package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100) + 1 //1-100

	num := 0
	for {
		fmt.Scan(&num)
		if num > value {
			fmt.Println("您输入的数太大了")
		} else if num < value {
			fmt.Println("您输入的数太小了")
		} else {
			fmt.Println("恭喜您猜对了")
			break;
		}
	}

}
