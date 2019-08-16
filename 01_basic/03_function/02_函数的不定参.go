package main

import "fmt"

func cacl(arr ... int) {
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0:2])
	fmt.Println(arr[0:])
	fmt.Println(arr[:])
	fmt.Println(arr[:2])


}
func main() {
	a, b, c := 1, 2, 3
	cacl(a, b, c, 4, 5, 6)
}
