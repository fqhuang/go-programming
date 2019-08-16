package main

import "fmt"

//阶乘
var sum = 1
func jiecheng(num int){
	if(num == 1){
		return
	}
	//sum *= num //5*4*3*2
	//fmt.Println(num)
	jiecheng(num -1)
	sum *= num //2*3*4*5
	fmt.Println(num)

}
func main() {
	a := 5
	jiecheng(a)
}
