package main

import "fmt"

func BubbleSort1(arr []int) {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("count", count)
}

//冒泡排序优化
func BubbleSort2(arr []int) {
	count := 0
	flg := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flg = true
			}
		}
		if !flg {
			fmt.Println("count", count)
			return
		} else {
			flg = false
		}
	}
}
func main() {
	arr := []int{1, 3, 5, 6, 3, 4, 5, 6, 3}
	//BubbleSort1(arr)
	BubbleSort2(arr)
	fmt.Println(arr)
}
