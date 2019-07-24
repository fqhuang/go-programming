package main

import "fmt"

func SelectSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		index := 0
		for j := 1; j < len(data)-i; j++ {
			if data[j] > data[index] {
				index = j
			}
		}
		//将最大值放在每行的末尾
		data[len(data)-1-i], data[index] = data[index], data[len(data)-1-i]
	}
}

func main() {
	arr := []int{1, 3, 5, 6, 3, 2, 4, 5, 6, 3, 9, 3, 5, 7, 9, 10, 0}
	SelectSort(arr)
	fmt.Println(arr)
}
