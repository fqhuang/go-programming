package main

import "fmt"

func main() {
	//定义并初始化map
	//map中的key必须支持== != 操作  不能使用函数 切片 map
	//在初始化中键必须是唯一的
	m := map[int]string{101: "周杰", 109: "周杰伦", 666: "张三", 102: "张三丰"}

	//重置指定key中的值
	m[101] = "马云"

	//fmt.Println(m)

	//打印值  在使用中只能通过key找到value  反之不可以
	fmt.Println(m[101])
}
