package main

/*
代码用途说明：
	演示Golang interface和nil使用时容易掉的坑

*/

import "fmt"

type IPeople interface {
	hello()
}
type People struct {
}

func (p *People) hello() {
	fmt.Println("github.com/meetbetter")
}

//错误：将nil的people给空接口后接口就不为nil,因为interface中的value为nil但type不为nil

func errFunc() *People {
	var p *People

	return p
}

//正确处理返回nil给接口的方法,返回时go就确定了接口是不是nil
func rightFunc() IPeople {
	var p *People

	return p
}
func main() {

	if errFunc() == nil {

		fmt.Println("对了哦，外部接收到也是nil")
	} else {

		fmt.Println("错了咦，外部接收到不是nil哦")

	}

	if rightFunc() == nil {

		fmt.Println("对了哦，外部接收到也是nil")
	} else {

		fmt.Println("错了咦，外部接收到不是nil哦")

	}

}
