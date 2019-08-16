package main

import "fmt"

func f1() {

	i := 0
	defer func(i int) {
		fmt.Println("defer func: ", i)
	}(i)

	i = 100

}

func f2() {

	i := 0
	defer func() {
		fmt.Println("defer func: ", i)
	}()

	i = 100

}
func main() {

	f1()
	f2()
}
