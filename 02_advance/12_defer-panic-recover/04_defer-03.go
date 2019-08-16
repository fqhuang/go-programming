package main

import "fmt"

/*
	演示defer的坑之"defer + print"：
	当defer后面接Print输出时会先将输出需要的值计算出来在压栈
*/

type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}

func (p *message) print1() string {
	//fmt.Println("print1 running")
	return p.content

}

func (p *message) print2() {
	//fmt.Println("print2 running")
	fmt.Println(p.content)

}

func test1() {
	m := &message{content: "Hello"}

	defer fmt.Println(m.print1())

	m.set("World")

	//fmt.Println("call print1")
}

func test2() {
	m := &message{content: "Hello"}

	defer m.print2()

	m.set("World")

	//fmt.Println("call print2")
}

func test3() string {
	fmt.Println("test3 Entry")

	return "github.com/meetbetter"
}

func test4() string {
	fmt.Println("test4 Entry")

	return "github.com/meetbetter"
}

func main() {
	test1()
	test2()

	//defer fmt.Println(test3())
	//defer test4()

	fmt.Println("main end")
}
