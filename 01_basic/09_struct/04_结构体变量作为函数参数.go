package main

import "fmt"

type students struct {
	id   int
	name string
	age  int
	sex  string
	addr string
}

//结构体变量作为函数参数
func test1(stu students) {
	//fmt.Println(stu)
	stu.name = ""
}

func main() {

	stu := students{1001, "张三", 18, "男", "中关村"}
	//结构体变量作为函数参数 值传递
	test1(stu)
	fmt.Println(stu)
}
