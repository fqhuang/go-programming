package main

import "fmt"

func main01() {
	var slice []int //空切片 指向内存地址编号为0空间
	fmt.Println("1:", cap(slice))
	//fmt.Printf("%p\n",slice)

	slice = append(slice, 1)
	fmt.Println("2:", cap(slice))

	slice = append(slice, 2)
	fmt.Println("3:", cap(slice))

	slice = append(slice, 2, 3, 4, 5, 7)
	fmt.Println(len(slice))
	fmt.Println("4:", cap(slice))

}

func main() {
	var slice []int
	for i := 0; i <= 100000; i++ {
		//如果没有超过1024 按照上一次倍数扩容 超过1024  按照上一次的1/4扩容
		slice = append(slice, i)
		if len(slice) == cap(slice)-1 {
			fmt.Println(cap(slice))
		}
	}
}
