package main

import "fmt"

func main01() {
	//var arr [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	//初始化部分元素 未初始化值为0
	//var arr [10]int=[10]int{1,2,3,4,5}
	//var arr [10]int = [10]int{0}
	//var arr [10]int = [10]int{1}
	//var arr [10]int

	//自动推导类型创建数组
	//arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
	//根据值的个数创建数组元素个数
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr) //[10]int
}

func main() {
	arr := [10]int{0}

	var arr1 [10]int = [10]int{1, 2, 3, 4, 5}
	//数组名  表示一个数组的集合  不能赋值一个变量
	//可以将一个数组赋值给另外一个数组 前提是相同的数据类型 相同的元素个数
	arr = arr1

	fmt.Println(arr)

}
