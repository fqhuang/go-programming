package main

import (
	"fmt"
	"login"
)


//全局变量  作用域  项目中所有文件 如果在同一包下 直接调用 如果在不用包下 需要使用包名.全局变量名 (首字母大写) 调用
var value int = 123
//函数的作用域  项目中的所有文件 如果在同一包下 直接调用 如果在不用包下 需要使用包名.函数名 (首字母大写) 调用

//多文件编译
//go build | run 源文件1.go 源文件2.go
func main() {
	sum:=Add(10,20)
	fmt.Println(sum)

	test()

	login.UserCreate()
	login.UserLogin()
	fmt.Println(login.Count)

}
