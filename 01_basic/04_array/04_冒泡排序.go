package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}

	fmt.Println("排序前：", arr)
	//外层控制行  表示执行次数
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列 表示比较次数
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素  升序  大于号  降序 小于号
			if arr[j] > arr[j+1] {
				//交换数据
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("排序后：", arr)
}
