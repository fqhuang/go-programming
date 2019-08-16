package main

import "fmt"

func main() {

	var p *int
	//new(数据类型)开辟数据类型对应的内存空间  并初始化
	p = new(int)
	//*p=123
	fmt.Println(*p)//打印指针指向地址中的数据
	fmt.Println(p) //打印指针变量里存储的地址
	//p = nil//将指针置为空 垃圾回收机制gc就会将内存空间回收
}
