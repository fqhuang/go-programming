package main

/*生成双色球*/

import (
	"fmt"
	"math/rand"
	"time"
)

//红球：1-33  选择6个不能重复  篮球1-16 选择1个  可以和红球重复
func main() {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		m := make(map[int]int)
		//随机红球
		for len(m) < 6 {
			m[rand.Intn(33)+1] = 0
		}

		//打印红球
		for k, _ := range m {
			fmt.Print(k, " ")
		}
		//打印篮球
		fmt.Println("+", rand.Intn(16)+1)
	}

}
