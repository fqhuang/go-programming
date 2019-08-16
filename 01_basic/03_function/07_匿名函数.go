package main

import "fmt"

//匿名函数重复用在同一个函数内部，可以实现简化函数和闭包
func main() {
	f := func(a int, b int) int{
		return a + b
	}

	sum1 := f(1, 2)
	fmt.Println(sum1)
	sum2 := f(2,3)
	fmt.Println(sum2)
}
