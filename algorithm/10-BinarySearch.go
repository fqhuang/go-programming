package main

import "fmt"

//二分查找 BinarySearch(数据，元素) 返回值为下标
func BinarySearch(arr []int, num int) int {

	start := 0               //定义起始下标
	end := len(arr) - 1      //定义结束下标
	mid := (start + end) / 2 //定义中间基准值

	for i := 0; i < len(arr); i++ {
		if num == arr[mid] { //中间基准值直接为查找值
			return mid //返回下标
		} else if num > arr[mid] { //查找值 比基准值大
			start = mid + 1 //查找右侧
		} else { //查找值 比基准值小
			end = mid - 1 //查找左侧
		}

		mid = (start + end) / 2 //根据新起始、结束下标，再次设置中间基准值
	}

	return -1 //没有找到
}

func main() {
	//有序数据
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	index := BinarySearch(arr, 9)
	fmt.Println(index)
}
