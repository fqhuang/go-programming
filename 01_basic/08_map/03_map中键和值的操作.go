package main

import "fmt"

func main01() {
	m := map[int]string{101: "aaa", 106: "bbb", 109: "ccc"}

	//判断key是否存在
	value, ok := m[111]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key不存在")
	}
	//fmt.Println(m[111])
}

func main() {
	m := map[int]string{101: "aaa", 106: "bbb", 109: "ccc"}

	//delete(map，key)
	delete(m, 111)
	fmt.Println(m)

}
