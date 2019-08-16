package main

import "fmt"




func main() {

	//计算圆的面积和周长
	//面积： PI*r*r
	//周长：2*PI*r

	//float32//单精度浮点型
	//float64//双精度浮点型
	//var PI float64=3.14159
	var PI float64
	PI=3.14159
	var r float64=2.5

	var s float64 = PI*r*r
	var l float64 = 2*PI*r

	fmt.Println("面积：",s)
	fmt.Println("周长：",l)
}