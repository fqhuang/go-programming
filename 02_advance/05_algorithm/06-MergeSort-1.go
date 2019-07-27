package main

import "fmt"

//将要比较的数据放到一个临时数组，对比临时数组中的数据再放回原数组
func merge1(arr []int, left, mid, right int) {
	oldArr := make([]int, right-left+1) //临时存放要排序的数据
	for i := left; i <= right; i++ {
		oldArr[i-left] = arr[i]
	}

	lIndex := left    //左侧数据的起点
	rIndex := mid + 1 //右侧数据的起点
	for i := left; i <= right; i++ {
		if lIndex > mid { //左侧数据已经全部归并完,只剩右边数据
			arr[i] = oldArr[rIndex-left]
			rIndex++
		} else if rIndex > right { //右侧数据已经全部归并完,只剩左边数据
			arr[i] = oldArr[lIndex-left]
			lIndex++
		} else {
			if oldArr[lIndex-left] < oldArr[rIndex-left] { //左侧数据小于右侧取右侧
				arr[i] = oldArr[lIndex-left]
				lIndex++
			} else { //否则右侧数据
				arr[i] = oldArr[rIndex-left]
				rIndex++
			}
		}
	}

}

//左右分成两个独立数组再对比，对比结果写回原数组
func merge2(arr []int, left, mid, right int) {
	//将arr分成左右两个数组
	arrLeft := make([]int, mid-left+1)
	arrRight := make([]int, right-mid)
	for i := left; i <= mid; i++ {
		arrLeft[i-left] = arr[i]
	}
	for i := mid + 1; i <= right; i++ {
		arrRight[i-(mid+1)] = arr[i]
	}

	i, j := 0, 0
	k := left
	for i < len(arrLeft) && j < len(arrRight) {

		if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}

		k++
	}

	if i >= len(arrLeft) {
		for j < len(arrRight) {
			arr[k] = arrRight[j]
			j++
			k++
		}
	}

	if j >= len(arrRight) {
		for i < len(arrLeft) {
			arr[k] = arrLeft[i]
			i++
			k++
		}
	}
}

//将两组数据对比后先赋给临时数组,再赋值给原数组-- 最好理解
func merge3(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)

	i, j := left, mid+1
	k := 0

	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}

	for j <= right {
		temp[k] = arr[j]
		k++
		j++
	}

	//赋值给原arr
	k = 0
	for left <= right {
		arr[left] = temp[k]
		k++
		left++
	}

}

//归并排序
func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2 //+1是避免对2取余导致最后一个数据丢失
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	if arr[mid] > arr[mid+1] { //当左边最后一个数据不大于右边第一个数据时不需要归并,减少不必要的操作
		//merge1(arr, left, mid, right)
		//merge2(arr, left, mid, right)
		merge3(arr, left, mid, right)
	}

}

//归并排序
func main() {
	arr := []int{0, 1, 9, 2, 8, 0, 3, 7, 4, 4, 6, 5, 5}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
