package main

import "fmt"

type Stu struct {
	name string
	age  int
	sex  string
	addr string
}

func main() {
	m := map[Stu]int{Stu{"张三", 18, "男", "中关村"}: 999}

	fmt.Println(m)
}
