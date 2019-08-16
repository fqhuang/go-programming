package main

import (
	"strconv"
	"fmt"
)

func main() {
	var slice []byte

	//将bool类型转字符切片
	slice = strconv.AppendBool(slice, false)
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendFloat(slice, 3.14, 'f', 6, 64)
	//将字符串转成字符切片 会有双引号
	slice = strconv.AppendQuote(slice, "hello")
	fmt.Println(string(slice))
}
