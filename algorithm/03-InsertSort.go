package main

import "fmt"

//插入排序法--使用向后挪移的方式操作有序数据
func InsertSort(arr []int) {
	//1 将arr分为有序数组和无序数组两部分
	for i := 1; i < len(arr); i++ { //arr[0]作第一个有序组
		//2 只有当无序数组的第一个数据arr[i]小于有序数组的最后一个数据arr[i-1]时,
		// 需要将该数据添加到有序数组
		if arr[i] < arr[i-1] {
			//3 循环判断要插入的数据应该插在有序数组的哪个位置
			tmp := arr[i] //备份要插入的数据,避免数据后移时被覆盖
			j := i - 1

			for j >= 0 && arr[j] > tmp { //循环和有序数组比较,确定插入位置
				arr[j+1] = arr[j] //有序数组中比要插入的数据大的一次向后挪移
				j--
			}
			arr[j+1] = tmp //j+1
		}
	}
}

//插入排序法--容易理解版--配合冒泡法进行有序数据的插入
func InsertSort2(arr []int) {

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			//tmp := arr[i]
			//j := i - 1
			for j := i - 1; j >= 0; j-- {
				if arr[j+1] < arr[j] {
					arr[j+1], arr[j] = arr[j], arr[j+1]
				}

			}
		}
	}
}

func main() {
	arr := []int{1, 3, 5, 6, 3, 4, 5, 6, 3}
	//InsertSort(arr)
	InsertSort2(arr)
	fmt.Println(arr)
}
