package main

import "fmt"

func main01() {
	//println表示输出数据后换行
	//fmt.Println("hello")
	//fmt.Println("world")
	//转义字符 \n 表示换行
	fmt.Print("hello\n")
	fmt.Print("world")
}

func main02(){
	a:=10
	b:=123456
	//格式化输出
	//占位符  %d 表示输出一个十进制整型数据
	//%nd  表示数据有n位 不满足n位左面补空格
	//%-nd 表示数据有n位 不满足n位右面补空格
	//%0nd 表示数据有n位 不满足n位左面补0
	//如果数据超出了n位正常输出
	fmt.Printf("==%030d==%5d\n",a,b);
}


func main() {
	a:=3.141595863

	//%f 占位符 表示输出一个浮点型数据
	//默认保留六位小数 会对第七位进行四舍五入
	//fmt.Printf("%f",a)
	//%.nf 小数点保留n位  会对n+1位进行四舍五入
	//fmt.Printf("%.3f",a)
	//%m.nf m表示整体位数 小数点算作一位  n 小数点保留n位  会对n+1位进行四舍五入
	fmt.Printf("==%6.3f==",a)


}