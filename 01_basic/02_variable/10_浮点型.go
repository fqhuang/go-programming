package main

import "fmt"

func main() {
	var a float32 = 3.14//单精度浮点型
	var b float64 = 3.14//双精度浮点型
	c:=3.14//float64

	//双精度浮点型比单精度浮点型更精准  建议采用float64 代替 float32
	fmt.Println(a,b)
	fmt.Printf("%.20f\n",a)
	fmt.Printf("%.20f\n",b)

	fmt.Printf("%T",c)
}
