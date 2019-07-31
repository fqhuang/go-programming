package main

import "fmt"

func main0101() {
	const count int = 10
	//数组中的元素个数不能 是变量 必须是一个常量 或常量表达式
	//数组的长度不能扩展
	var arr [10]int
	fmt.Println(arr)
}


func main0102(){
	//切片的定义 var 切片名 []数据类型
	var slice []int//空切片
	//slice[0]=123//下标越界 err

	//为切片添加数据用append
	//切片名=append(切片名,元素1，元素2)
	slice=append(slice,123,456,789)
	//slice=append(slice,456,678)
	fmt.Println(slice)
}


func main(){
	//make(切片，长度)
	var slice []int =make([]int,10)
	fmt.Println(slice)
	//打印切片元素个数
	fmt.Println(len(slice))//长度
	fmt.Println(cap(slice))//容量
	//在数据末尾添加数据  追加
	slice=append(slice,123,456)
	fmt.Println(len(slice))//长度
	//容量扩充为上一次的2倍数
	fmt.Println(cap(slice))//容量
}