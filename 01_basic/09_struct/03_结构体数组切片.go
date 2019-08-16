package main

import "fmt"

type Students struct {
	id   int
	name string
	age  int
	sex  string
	addr string
}

//结构体数组作为函数参数
func BubbleSort(arr [5]Students) [5]Students {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//结构体变量只能使用== !=进行比较操作
			//比较两个结构体成员
			if arr[j].age < arr[j+1].age {
				//交换结构体变量
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return arr
}

//结构体切片作为函数参数
func BubbleSort1(slice []Students) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			//结构体变量只能使用== !=进行比较操作
			//比较两个结构体成员
			if slice[j].age < slice[j+1].age {
				//交换结构体变量
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}

	}
}

func main() {

	//结构体切片
	var slice []Students = []Students{
		Students{1001, "杨过", 18, "男", "终南山下"},
		Students{1002, "小龙女", 22, "女", "终南山下"},
		Students{1003, "尹志平", 20, "男", "终南山上"},
		Students{1004, "洪七公", 58, "男", "丐帮"},
		Students{1005, "李莫愁", 38, "女", "终南山下 "},
	}

	//为切片添加元素
	slice = append(slice, Students{1006, "欧阳锋", 52, "男", "白拓山"})

	BubbleSort1(slice)
	fmt.Println(slice)
	//计算切片的长度和容量
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
