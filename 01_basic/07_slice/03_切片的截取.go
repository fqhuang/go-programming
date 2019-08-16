package main

import "fmt"

func main01() {
	slice := []int{1, 2, 3, 4, 5}
	//切片名 本身就是一个地址  存储三个值  指针  长度  容量·
	//fmt.Printf("%p\n",slice)

	//计算数据类型在内存中占的字节大小
	//fmt.Println(unsafe.Sizeof(slice))
	//用一个切片为另外一个切片赋值
	s := slice //浅拷贝

	//改变slice 的值  影响s的值
	slice[0] = 123
	fmt.Println(s)
	fmt.Println(slice)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", slice)
}

func main02() {
	slice := []int{1, 2, 3, 4, 5}

	s := make([]int, 5)

	//copy(目标，源)
	copy(s, slice) //深拷贝 拷贝的是内容

	//修改slice
	slice[0] = 123
	fmt.Println(s)
	fmt.Println(slice)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", slice)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//切片截取还是源切片的一部分 修改会影响源切片
	//切片[起始位置：结束位置]  不包含结束位置
	//s := slice[0:5]
	//fmt.Printf("%p\n", s)
	//fmt.Printf("%p\n", slice)
	//fmt.Println(s)

	//切片[起始位置:]  从起始位置到结尾
	//s:=slice[2:]
	//fmt.Println(s)
	//切片[:结束位置]  从0开始到结束位置 不包含结束位置
	//s := slice[:2]
	//fmt.Println(s)

	//切片[:]  从0到结束位置
	//s:=slice[:]
	//fmt.Println(s)

	//切片[起始位置：结束位置：最大容量]  cap = 最大容量-起始位置  最大容量大于结束位置
	s := slice[2:5:6]
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println(s)
}
