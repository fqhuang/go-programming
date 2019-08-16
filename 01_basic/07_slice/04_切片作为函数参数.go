package main

import "fmt"

//切片作为函数参数
func BubbleSort(slice []int) {
	//fmt.Println(slice)
	//fmt.Printf("%p\n",slice)

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
func main() {
	slice := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}
	//地址传递  引用传递  形参和实参指向相同的地址单元  形参可以修改实参的值
	BubbleSort(slice)
	//fmt.Printf("%p\n", slice)

	//fmt.Printf("%T\n", slice)
	fmt.Println(slice)
}
