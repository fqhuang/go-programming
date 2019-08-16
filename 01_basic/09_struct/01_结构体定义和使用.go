package main

import "fmt"

//结构体定义
/*
type 结构体名称 struct{
	结构体成员
}
*/
type Student struct {
	//结构体成员
	id   int
	name string
	age  int
	sex  string
	addr string
}

func main() {

	//结构体变量的定义
	//var 变量 结构体数据类型
	var stu Student
	//为结构体成员赋值
	//结构体变量名.成员名
	stu.id = 1001
	stu.age = 18
	stu.name = "张三"
	stu.addr = "中关村"
	stu.sex = "女装大佬"

	fmt.Println(stu)
	//打印成员信息
	fmt.Println(stu.name)
	fmt.Println(stu.age)
	fmt.Println(stu.sex)

}
