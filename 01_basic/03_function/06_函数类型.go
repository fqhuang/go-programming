package main

import "fmt"

//type FUNCTYPE func(int, int) (int, int)
type FUNCTYPE = func(int, int) (int, int) //加不加 = 的打印出来的数据类型有区别

func demo040901(a int, b int) (sum int, sub int) {
	fmt.Println("helloworld")
	return
}
func main01() {
	var f FUNCTYPE
	f = demo040901
	f(1,2)
	fmt.Printf("%T\n", f)
}


//type INT int
type INT = int

func main(){
	var a INT = 1
	var b int = 2
	//type INT int没有等号则不是同样一个数据类型，不能直接运算invalid operation: a + b (mismatched types INT and int)
	c := a + b

	fmt.Println(c)
}