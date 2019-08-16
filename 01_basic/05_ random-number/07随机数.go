package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	//创建随机数种子  混淆  伪随机数
	rand.Seed(time.Now().UnixNano())
	//打印系统时间
	//fmt.Println(time.Now().Second())

	//生成随机数
	//value:=rand.Intn(10)
	//fmt.Println(value)

	for i:=0;i<10;i++{
		fmt.Println(rand.Intn(100))
	}

}
