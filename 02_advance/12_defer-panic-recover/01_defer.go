package main

import (
	"fmt"
	"os"
)

func f1() {
	defer fmt.Println("1111")

	fmt.Println(2222)
}

func f2() {
	i := 10
	defer func() {
		i = 100
	}()

	fmt.Println("i: ", i)
}

func f3() {
	defer fmt.Println("1111")
	defer fmt.Println("2222")
	defer fmt.Println("3333")
	defer fmt.Println("4444")
}

func f4() error {
	f, err := os.Open("a.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
func main() {
	f1()
	f2()
	f3()

	f4()

}
