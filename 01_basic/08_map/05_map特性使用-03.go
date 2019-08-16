package main

/*
	利用map在切片中查找和为target的两个元素的下标
*/

import "fmt"

func twoSum(nums []int, tgt int) []int {

	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int)

	for i, v := range nums {
		sub := tgt - v
		m[sub] = i

	}

	for j := 0; j < len(nums); j++ {

		if _, ok := m[nums[j]]; ok {

			if j != m[nums[j]] { //防止同一下标被统计两次
				return []int{m[nums[j]], j}
			}

		}
	}

	return nil
}

func main() {
	nums := []int{1, 2, 3, 8, 9}
	fmt.Println(twoSum(nums, 9))
}
