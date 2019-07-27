package main

import "fmt"

func main() {
	//变量的声明
	//var a int
	//变量赋值
	//a=10

	//自动推导类型  是根据变量后面的值进行类型推导
	a := 10//int
	//var a int
	//a := 10//err

	b := 3.14//float64

	fmt.Println(a)
	fmt.Println(b)
}

//显示版本 go version
//编译Go语言程序
//go build 源文件.go
//编译并运行Go语言程序
//go run 源文件
