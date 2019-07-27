package main

import "fmt"

func main() {
	//int表示有符号的整型数据
	//uint表示无符号整型数据
	//根据大小 8 16 32 64

	//int 类型会根据系统类型不同 创建的字节大小也不同
	var a uint = 10
	//fmt.Println(a)
	fmt.Printf("%d",a)
}
