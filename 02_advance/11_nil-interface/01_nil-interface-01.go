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
func errFunc1(in int) *People {
	if in == 0 {
		fmt.Println("importantFunc返回了一个nil")
		return nil
	} else {
		fmt.Println("importantFunc返回了一个非nil值")
		return &People{}
	}

}

//正确处理返回nil给接口的方法,返回时go就确定了接口是不是nil
func rightFunc(in int) IPeople {
	if in == 0 {
		fmt.Println("importantFunc返回了一个nil")
		return nil
	} else {
		fmt.Println("importantFunc返回了一个非nil值")
		return &People{}
	}

}

func main() {
	var i IPeople

	in := 0
	//i = rightFunc(in)
	i = errFunc1(in)

	if i == nil {

		fmt.Println("哈，外部接收到也是nil")
	} else {

		fmt.Println("咦，外部接收到不是nil哦")
		fmt.Printf("%v, %T\n", i, i)
	}

}
