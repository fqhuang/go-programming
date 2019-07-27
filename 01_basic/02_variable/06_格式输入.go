package main

import "fmt"

func main0601() {

	//如果变量没有初始值 默认值为0
	var a int

	//& 取地址运算符
	//fmt.Println(&a)
	//通过键盘赋值
	//fmt.Scan 阻塞式请求
	fmt.Scan(&a)
	fmt.Println(a)
}

func main0602(){
	//var a int
	//var b int
	var a, b int
	//可以使用空格或换行表示一个数据接收结束
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
}

func main(){
	var a int
	var b int
	//表示接收前3位数据
	fmt.Scanf("%3d%d",&a,&b)
	fmt.Println(a,b)
}
