package main

import "fmt"

func myAtoi(str string) int {

	var num int = 0
	i := 0
	flg := true
	if str[0] == '-' {
		flg = false
		i++
	}
	for ; i < len(str); i++ {
		if '0' <= str[i] && str[i] <= '9' {
			fmt.Println(str[i])
			fmt.Printf("%T\n", str[i])

			num *= 10
			num += int(str[i] - '0')
		}
	}
	if !flg {
		return 0 - num
	}
	return num
}

func main() {
	str := "-0123405"

	num := myAtoi(str)
	fmt.Println(num)
}
