package main

import "fmt"

func demo1(arr ...int){
	fmt.Println(arr)
}

func demo2(arr ...int){
	demo1(arr ...)
	//demo1(arr[0], arr[1], arr[2])
	//demo1(arr[1:] ...)
	//demo1(arr[:] ...)
	//demo1(arr[:3] ...)


}
func main() {
	demo2(1, 2, 3, 4, 5, 6)
}
