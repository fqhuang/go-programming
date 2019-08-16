package main

import "fmt"

func main01() {

	//字符类型
	var a byte
	//使用单引号引起来成为字符
	a = 'A'
	//ASCII 码 格式打印数据
	fmt.Println(a)//97

	//%c 是一个占位符表示输出一个字符类型数据
	fmt.Printf("%c",a)
}

func main(){
	var ch byte//uint8
	//fmt.Scan(&ch)//err
	//如果接收字符 必须使用fmt.Scanf
	fmt.Scanf("%c",&ch)
	fmt.Printf("%c",ch+32)
	fmt.Printf("%T",ch)
}
