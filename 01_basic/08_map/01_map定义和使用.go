package main

import "fmt"

func main() {
	//map定义格式
	//var m map[键类型]值类型  map[key]value
	//var m map[int]string = make(map[int]string)
	m := make(map[int]string)
	m[101] = "张三"
	m[109] = "李四"
	m[666] = "王五"

	//map的存储是无序的
	fmt.Println(m)
	//范围遍历
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
}
