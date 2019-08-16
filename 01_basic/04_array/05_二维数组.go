package main

import "fmt"

func main1001() {
	//var 数组名 [元素个数]数据类型
	//var 数组名 [行个数][列个数]数据类型

	//var arr [2][3]int=[2][3]int{{1,2,3},{4,5,6}}
	var arr [2][3]int = [2][3]int{{1, 2}, {4}}
	//为数组元素赋值  数组名[行下标][列下标]=值

	//arr[1][1] = 123
	//arr[1][2] = 456

	//fmt.Println(arr)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j], "  ")
		}
		fmt.Println()
	}

}

func main1002() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//len(二维数组名)  一个二维数组行数
	//fmt.Println(len(arr))
	//len(二维数组名[0])  一个二维数组列数
	//fmt.Println(len(arr[0]))

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], "  ")
		}
		fmt.Println()
	}
}

func main() {
	var arr [2][3][4]int = [2][3][4]int{
		{
			{1, 2, 3, 4},
			{2, 3, 4, 5},
			{3, 4, 5, 6},
		},
		{
			{4, 5, 6, 7},
			{5, 6, 7, 8},
			{6, 7, 8, 9},
		},
	}

	//fmt.Println(arr)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				fmt.Print(arr[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println("-------------")
	}
}
