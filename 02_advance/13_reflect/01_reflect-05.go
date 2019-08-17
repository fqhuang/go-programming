package main

/*
通过reflect进行方法调用
*/

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello1(name string, age int) {
	fmt.Println("Hello1 name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) Hello2() {
	fmt.Println("Hello2")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.Hello直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动methodValue.Call

func main() {
	user := User{1, "原作者Allen.Wu", 25}

	fmt.Println(reflect.TypeOf(user).Name())

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("Hello1")
	args := []reflect.Value{reflect.ValueOf("github.com/meetbetter测试"), reflect.ValueOf(18)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = getValue.MethodByName("Hello2")
	//args = make([]reflect.Value, 0)
	args = []reflect.Value{}
	methodValue.Call(args)
}
