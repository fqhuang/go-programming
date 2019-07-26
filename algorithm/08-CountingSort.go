package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 计数统计排序
func CountingSort(s []int) {
	// 创建 map 统计 1-999 每个数出现的 次数.
	m := make(map[int]int)

	// 遍历 切片,统计各个数出现的次数,使用 map 记录
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	fmt.Println("\n", m)

	// 按序借助map 输出所有数据
	for i := 0; i < 1000; i++ {
		// 内层循环按 每个数,出现的次数打印
		for j := 0; j < m[i]; j++ {

			fmt.Print(i, " ")
		}
	}

}
func main() {
	s := make([]int, 0)

	rand.Seed(time.Now().UnixNano()) // 播种随机数种子

	// 循环生成100000个随机数
	for i := 0; i < 100000; i++ {
		s = append(s, rand.Intn(1000))
	}
	fmt.Println(s)

	CountingSort(s)
}
