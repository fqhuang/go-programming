package main

import "fmt"

func test1(a int, b int) int{
	return a + b
}
func test2(a int, b int) int{
	return test1(a, b)
}

func main() {
	a, b := 10, 20

	sum := test2(a, b)

	fmt.Println(sum)
}
