package main

import (
	"math/rand"
	"time"
	"fmt"
)

//通过函数实现冒泡排序
//数组作为函数参数  数组作为函数返回值
//函数定义
func BubbleSort(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	//fmt.Println("函数：",arr)
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := [10]int{} //[10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("排序前：", arr)

	//函数调用  数组作为函数参数 值传递  形参和实参是不同的存储单元 改变形参不影响实参的值
	arr = BubbleSort(arr)
	fmt.Println("排序后：", arr)

}
