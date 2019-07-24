package main

import "fmt"

func QuickSort(arr []int, left, right int) {

	//1. 设置基准值和基准值所在下标
	pivot := arr[left]  //基准值
	index := left       //基准下标
	i, j := left, right //本次比较开始时的左右下标

	//2. 将本组数据小于pivot的移动到左侧,大于pivot的移动到pivot右侧
	for i <= j { //循环直到i，j相遇，即一组比较完成
		//2.1 pivot和右侧比较,有小于pivot的则和pivot交换位置
		for j >= index && arr[j] >= pivot { //左移直到找到比基准值小的数据或者i,j相遇
			j--
		}
		if j >= index { //i>=index成立说明是找到了比pivot小的数据
			arr[index] = arr[j] //将小值移动到pivot所在位置
			index = j           //将index指向比pivot小的位置
		}

		//2.2 pivot和左侧比较,有大于pivot的则和pivot交换位置
		for i <= index && arr[i] <= pivot { //右移直到找到比基准值大的数据或者i,j相遇
			i++
		}
		if i <= index { //i<=index成立说明arr[i] <= pivot不成立,即左侧找到比pivot大的数据
			arr[index] = arr[i] //将大值移动到当前index指向的位置
			index = i           //index指向此处
		}
	}

	//3. 此时基准值的位置永久固定，左侧均小于pivot，右侧均大于pivot
	arr[index] = pivot //比较完成后pivot应该在index处，左侧比pivot小，右侧比pivot大

	//4. 递归进入两个新的分组执行比较和分组，直到全部比较完成
	if index-left > 1 { //只有两个数据以上才需要比较
		QuickSort(arr, left, index-1)
	}
	if right-index > 1 {
		QuickSort(arr, index+1, right)
	}

}

func main() {
	arr := []int{1, 3, 5, 6, 3, 4, 59, 100, 200, 16, 0, 7, 10, 1, 6, 3}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
