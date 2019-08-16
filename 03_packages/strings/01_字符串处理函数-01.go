package main

import (
	"strings"
	"fmt"
)

func main01() {
	str1 := "github.com/meetbetter"

	str2 := "go rprogramming"
	//用作于模糊查找
	//b := strings.Contains(str1, str2)
	//fmt.Println(b)
	if strings.Contains(str1, str2) {
		fmt.Println("出现")
	} else {
		fmt.Println("未出现")
	}

}

func main0702() {
	//slice := []string{"1111", "2222", "3333", "4444"}
	//字符串拼接
	slice := []string{"7653", "9101", "6910", "9643"}
	str := strings.Join(slice, "")
	fmt.Println(str)
}

func main0703() {
	str := "github.com/meetbetter "
	ch := "meet"

	//查找一个字符串在另外一个字符串出现的位置 返回值为下标 未找到值为 -1
	index:=strings.Index(str,ch)
	fmt.Println(index)
}


func main(){
	str:="github.com/meetbetter "

	//字符串重复n次
	ch:=strings.Repeat(str,3)
	fmt.Println(ch)

}