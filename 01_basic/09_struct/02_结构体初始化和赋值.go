package main

import (
	"fmt"
	"unsafe"
)

//在同一个包下不能重名
type student struct {
	id   int    //8
	name string //16
	age  int    //8
	sex  string //16
	addr string //16
}

func main0801() {
	//结构体赋值
	stu := student{1001, "张三丰", 180, "男", "武当山"}
	stu.name = "张君宝"
	fmt.Println(stu)

	//计算字符串在内存中占的字节大小
	fmt.Println(unsafe.Sizeof(stu.name))
	//计算结构体变量在内存中占的大小
	fmt.Println(unsafe.Sizeof(stu))

	//打印结构体变量内存地址
	fmt.Printf("%p\n", &stu)
	fmt.Println(&stu.id)
	fmt.Println(&stu.name)
	fmt.Println(&stu.age)
	fmt.Println(&stu.sex)
	fmt.Println(&stu.addr)
}

func main() {
	//结构体赋值
	stu := student{1002, "郭襄", 18, "女", "峨眉山"}

	//var s student
	//s=stu
	s := stu

	s.name = "灭绝师太"

	fmt.Println(stu)
	fmt.Println(s)

	fmt.Println(&stu.name)
	fmt.Println(&s.name)
}
