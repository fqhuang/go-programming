package main

import "fmt"

/*使用map统计每个字符的出现次数*/
func main() {
	str := "sdhdjfhfiuewyjbjuewdjkshgjsgdgsgdfjgsdfjsugfsdfdfhgjf"
	m := make(map[byte]int)
	//m['s']++
	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])
		m[str[i]]++
	}

	//fmt.Println(m['s'])
	//fmt.Println(m)
	for k, v := range m {
		fmt.Printf("字符：%c  出现：%d次\n", k, v)
	}
}
