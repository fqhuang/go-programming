package main

import "fmt"

type StackNode struct {
	Data interface{}
	Next *StackNode
}

//创建栈
func CreateStack(data ...interface{}) *StackNode {
	if data == nil || len(data) == 0 {
		return nil
	}

	var pre *StackNode
	for _, v := range data {
		newNode := new(StackNode)
		newNode.Next = pre
		newNode.Data = v

		pre = newNode //倒着存
	}

	return pre
}

//打印栈
func PrintStack(s *StackNode) {
	if s == nil {
		return
	}

	for s != nil {
		fmt.Print(s.Data, " ")
		s = s.Next
	}

	fmt.Println()
}

//获取栈长度
func StackLength(s *StackNode) int {
	if s == nil {
		return -1
	}

	i := 0

	for s != nil {
		i++
		s = s.Next
	}

	return i
}

//压栈 -- 头插
func PushStack(s *StackNode, data interface{}) *StackNode {
	if s == nil {
		return nil
	}

	if data == nil {
		return s //s not nil
	}

	newNode := new(StackNode)
	newNode.Data = data
	newNode.Next = s

	return newNode

}

//出栈 -- 头删
func PopStack(s *StackNode) *StackNode {
	if s == nil {
		return nil
	}

	return s.Next
}

func main() {
	s := new(StackNode)
	s = CreateStack(1, 2, 3, 4, 5)
	PrintStack(s)
	fmt.Println("length: ", StackLength(s))

	s = PushStack(s, 100)
	PrintStack(s)
	fmt.Println("length: ", StackLength(s))

	s = PopStack(s)
	PrintStack(s)
	fmt.Println("length: ", StackLength(s))
}
