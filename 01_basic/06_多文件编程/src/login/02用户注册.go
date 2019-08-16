package login

import "fmt"

//如果函数首字母大写表示可以在外部调用该函数
func UserCreate(){
	fmt.Println("用户注册成功")
}

//如果函数名首字母小写 表示本包的私有函数 只能在本包内部使用
//func userCreate(){
//	fmt.Println("用户注册成功")
//}