package main

import "fmt"

//将map作为函数参数
func test(m map[int]string) {
	//m[101] = "aaa"
	//fmt.Println(m)
	//m[188]="bbb"
	delete(m,188)
}

func main() {

	m := make(map[int]string)
	m[101] = "aaa"
	m[102] = "bbb"
	m[105] = "ccc"
	m[188] = "ddd"
	//map名是一个地址 地址传递  引用传递  形参可以改变实参的值
	test(m)
	//map数据类型
	//fmt.Printf("%T\n", m)

	fmt.Println(m)
}
