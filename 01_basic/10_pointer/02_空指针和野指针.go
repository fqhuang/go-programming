package main

import "fmt"

func main() {

	//空指针  指针变量指向内存地址空间编号为0
	var p *int //默认值为nil  值为0  0x0

	//内存地址编号为0的空间为系统占用 不允许用户读写操作
	*p = 123 //panic
	//fmt.Println(p)
	fmt.Println(*p) //panic

}

func main02() {
	//野指针 指针变量指向了一个未知的空间 在go语言中 不允许操作野指针
	//var p *int = *int(0x25186310)//err
}
