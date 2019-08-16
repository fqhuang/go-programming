package main

import "fmt"

func f1() {
	fmt.Println("github.com/meetbetter/go-programming")
	panic("boom")
	fmt.Println("f1 end")
}
func main() {
	fmt.Println("main Entry")
	f1()
	fmt.Println("main End")

}
