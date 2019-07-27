package main

import "fmt"

func main0401() {
	//变量的定义
	//var a int =10
	//var b int =20
	//var c int =30

	//自动推导类型
	//a:=10
	//b:=20
	//c:=30

	//多重赋值
	//a,b,c:=10,20,30
	//""引起来的成为字符串类型 string
	a, b, c := 10, 3.14, "hello"

	//快捷注释 ctrl+/
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	fmt.Println(a, b, c)

}

func main0402() {
	//_表示匿名变量
	a, _, _ := 10, 20, 30
	fmt.Println(a)
	//不能打印匿名变量的值
	//fmt.Println(_)
}

func main() {

	//var a int =10
	//var b int =20
	//a:=10
	//b:=20
	//通过第三方变量进行数据交换
	//temp:=a
	//var temp int = a//10
	//a=b//20
	//b=temp//10

	//a:=10
	//b:=20
	//
	//temp:=a
	//a=b
	//b=temp

	//使用多重赋值交换变量
	//a:=10
	//b:=20
	//
	//a,b = b,a

	//使用变量计算交换数据
	a := 10
	b := 20
	a = a + b //30
	b = a - b //10
	a = a - b //20

	fmt.Println(a)
	fmt.Println(b)

}
