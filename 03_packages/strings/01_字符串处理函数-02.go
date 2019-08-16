package main

import (
	"fmt"
	"strings"
)

func main01() {
	//str:="hello"
	str := "github.com/meetbetter"
	//字符串替换
	//ch:=strings.Replace(str,"l","mmm",2)
	ch := strings.Replace(str, "meetbetter", "**", -1)
	fmt.Println(ch)
}

func main02() {

	str := "github.com/meetbetter"

	//字符串切割
	slice := strings.Split(str, "t")

	fmt.Println(slice)
}

func main03() {
	str := "====aaaaaaaaa==========ah=e l l o==aaaaaaa========="

	//去掉字符串首位的指定字符
	ch := strings.Trim(str, "=ab")
	fmt.Println(ch)

}

func main() {
	str := "are u ok"
	//去掉空格  并把字符串转成 字符切片
	slice := strings.Fields(str)
	for _, v := range slice {
		fmt.Println(v)
	}

}
