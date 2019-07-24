package main

import "fmt"

func mergeArr(arr1 []int, arr2 []int) []int {
	temp := make([]int, 0)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] { //去重,i,j同时后移
			if len(temp) == 0 {
				temp = append(temp, arr1[i])
			} else if len(temp) >= 1 && arr1[i] != temp[len(temp)-1] {
				temp = append(temp, arr1[i])
			}
			i++
			j++
		} else if arr1[i] < arr2[j] {
			if len(temp) == 0 {
				temp = append(temp, arr1[i])
			} else if len(temp) >= 1 && arr1[i] != temp[len(temp)-1] {
				temp = append(temp, arr1[i])
			}
			i++
		} else {
			if len(temp) == 0 {
				temp = append(temp, arr2[j])
			} else if len(temp) >= 1 && arr2[j] != temp[len(temp)-1] {
				temp = append(temp, arr2[j])
			}
			j++

		}
	}

	for i < len(arr1) {
		if len(temp) == 0 {
			temp = append(temp, arr1[i])
		} else if len(temp) >= 1 && arr1[i] != temp[len(temp)-1] {
			temp = append(temp, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		if len(temp) == 0 {
			temp = append(temp, arr2[j])
		} else if len(temp) >= 1 && arr2[j] != temp[len(temp)-1] {
			temp = append(temp, arr2[j])
		}
		j++
	}

	return temp
}

//将两组数据对比后先赋给临时数组,再赋值给原数组-- 最好理解
func merge3(arr []int, left, mid, right int) int {
	temp := make([]int, 0)

	i, j := left, mid+1

	for i <= mid && j <= right {
		if arr[i] == arr[j] {
			if len(temp) == 0 {
				temp = append(temp, arr[i])
			} else if len(temp) >= 1 && arr[i] != temp[len(temp)-1] { //避免一组数据有两个以上的重复数据
				temp = append(temp, arr[i])
			}
			i++
			j++ //去重
		} else if arr[i] < arr[j] {
			if len(temp) == 0 {
				temp = append(temp, arr[i])
			} else if len(temp) >= 1 && arr[i] != temp[len(temp)-1] {
				temp = append(temp, arr[i])
			}
			i++
		} else {
			if len(temp) == 0 {
				temp = append(temp, arr[j])
			} else if len(temp) >= 1 && arr[j] != temp[len(temp)-1] {
				temp = append(temp, arr[j])
			}
			j++
		}
	}

	for i <= mid {
		if len(temp) == 0 {
			temp = append(temp, arr[i])
		} else if len(temp) >= 1 && arr[i] != temp[len(temp)-1] {
			temp = append(temp, arr[i])
		}
		i++
	}

	for j <= right {
		if len(temp) == 0 {
			temp = append(temp, arr[j])
		} else if len(temp) >= 1 && arr[j] != temp[len(temp)-1] {
			temp = append(temp, arr[j])
		}
		j++
	}

	k := 0
	for k < len(temp) {
		arr[left] = temp[k]
		k++
		left++
	}

	return len(temp)
}

//归并排序
func MergeSort(arr []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2 //+1是避免对2取余导致最后一个数据丢失
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	length := 0
	if arr[mid] > arr[mid+1] { //当左边最后一个数据不大于右边第一个数据时不需要归并,减少不必要的操作
		//merge1(arr, left, mid, right)
		//merge2(arr, left, mid, right)
		length = merge3(arr, left, mid, right)
	}

	return length //记录本组数据归并去重后的数据长度
}

//有BUG!!!不能正确截取去重后的数据。
//合并(去重)两组数据并排序
func main() {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 9, 8, 7, 6, 5, 4, 3, 2, 1, 12, 12, 0}
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//两组数据用归并排序并去除本组数据的重复数据
	length1 := MergeSort(arr1, 0, len(arr1)-1)
	length2 := MergeSort(arr2, 0, len(arr2)-1)
	//两组数据合并去重
	arr3 := mergeArr(arr1[:length1], arr2[:length2]) //截取有效数据,因arr1去重后有效数据减少但是main中arr1长度未变
	fmt.Println(arr3)
}
