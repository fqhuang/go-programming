package main

import (
	"strconv"
	"fmt"
)

func main0901() {
	//将其他类型转成字符串

	//布尔类型转成字符串
	//b:=false
	//str:=strconv.FormatBool(b)
	////fmt.Println(str)
	//fmt.Printf("%q\n",str)

	//整型转成字符串
	//str:=strconv.FormatInt(123456,10)
	//fmt.Println(str)

	//str := strconv.Itoa(123456)
	//fmt.Println(str)

	//浮点型转成字符串
	//strconv.FormatFloat(浮点型数据,'f',小数保留位数,处理位数)
	str := strconv.FormatFloat(3.1415999, 'f', 5, 64)
	fmt.Println(str)

}

func main() {
	//将字符串转成其他类型

	//将字符串转成布尔类型
	//b, err := strconv.ParseBool("trrr")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(b)

	//将字符串转成整型
	//v,err:=strconv.ParseInt("123",2,64)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(v)

	//v,_:=strconv.Atoi("123")
	//fmt.Println(v)

	f,_:=strconv.ParseFloat("1.23456",64)
	fmt.Println(f)

}
