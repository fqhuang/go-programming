package main

import "fmt"

func main01() {
	//var slice []int
	//var slice []int = make([]int,10)

	//make([]数据类型,长度,容量)  长度 小于等于 容量
	var slice []int = make([]int, 10, 11)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func main02() {
	//自动推导类型创建切片 长度和容量相同
	//slice := []int{1, 2, 3, 4, 5}
	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//fmt.Printf("%T\n", slice)

	//slice := make([]int, 10)
	slice := make([]int, 10, 20)
	fmt.Println(slice)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//slice=append(slice,11)
	//打印切片所有元素值  在打印时 使用len  不能使用cap
	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}

	//范围遍历 i 表示下标 v 表示值  如果不想接收 用匿名变量表示
	for i, v := range slice {
		//fmt.Println("下标：", i, "值：", v)
		fmt.Println(i, v)
	}
}
