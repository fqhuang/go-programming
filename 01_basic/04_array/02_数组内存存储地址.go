package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//数组名的地址 就是数组首元素地址
	fmt.Printf("%p\n",&arr)
	for i := 0; i < len(arr); i++ {
		//取出数组的元素的内存地址
		fmt.Println(&arr[i])
		//fmt.Printf("%p\n",&arr[i])
	}
}
