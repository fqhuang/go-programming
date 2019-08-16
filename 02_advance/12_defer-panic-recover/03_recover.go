package main

import "fmt"

func f1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("github.com/meetbetter/go-programming")

	fmt.Println("看不到我哦")
}

func main() {
	f1()

	fmt.Println("main End")
}
