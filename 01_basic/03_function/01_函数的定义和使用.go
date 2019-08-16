package main

import "fmt"

/*
func 函数名（函数参数列表）函数返回值类型｛
	函数体
	return 返回值
｝
*/


func sum_sub(a int, b int)(sum int, sub int){
	sum = a + b
	sub = a - b

	return
}

func main() {
	a, b := 20, 10

	sum, sub := sum_sub(a, b)

	fmt.Println(sum, sub)
}
