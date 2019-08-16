package main

import "fmt"

func main() {
	var a int = 10

	//fmt.Println(&a)
	//定义指针变量
	//var 指针变量名 *数据类型
	//var p *int = &a
	p := &a
	//fmt.Printf("%T\n",p)

	//通过间接修改变量的值
	*p = 123

	fmt.Println(a)

	fmt.Println(p)
	fmt.Println(&a)
}
