package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//起始下标
	start := 0
	//最大值下标
	end := len(arr) - 1
	for start < end{
		//交换数据
		arr[start], arr[end] = arr[end], arr[start]
		//改变坐标
		start++
		end--
	}
	fmt.Println(arr)
}
