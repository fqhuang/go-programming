package main

import "fmt"

func demo040801(){
	fmt.Println("demo040801")
}

func demo040802(a int, b int){
	fmt.Println("demo040802")
}

func demo040803(a int, b int) int{
	fmt.Println("demo040803")
	return a+b
}

func demo040804(a int, b int) (int, int){
	fmt.Println("demo040804")
	return a + b, a - b
}

func demo040805() (a int, b int){
	fmt.Println("demo040805")
	return a + b, a- b
}
func main() {
	fmt.Printf("%T\n", demo040801)
	fmt.Printf("%T\n", demo040802)
	fmt.Printf("%T\n", demo040803)
	fmt.Printf("%T\n", demo040804)
	fmt.Printf("%T\n", demo040805)

	//函数类型的变量
	var f func(int, int) (int, int)

	f = demo040804

	sum, sub := f(3, 2)

	fmt.Println(sum, sub)

}
