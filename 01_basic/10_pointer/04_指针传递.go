package main

import "fmt"

//func test(a int, b int) {
//	a, b = b, a
//}

//指针作为函数参数
func test(a *int, b *int) {

	*a, *b = *b, *a
}
func main() {
	a := 10
	b := 20
	//值传递  形参不能改变实参的值
	//test(a, b)
	//地址传递 引用传递  形参可以改变实参的值
	test(&a, &b)
	fmt.Println(a, b)

}
