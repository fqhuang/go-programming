package main

import "fmt"

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

//创建队列
func (node *QueueNode) Create(data ...interface{}) {
	if node == nil {
		return
	}

	for _, v := range data {
		newNode := new(QueueNode)
		newNode.Next = nil
		newNode.Data = v

		node.Next = newNode

		node = node.Next
	}
}

//获取队列长度
func (node *QueueNode) Length() int {
	if node == nil {
		return -1
	}

	i := 0
	node = node.Next
	for node != nil {
		i++
		node = node.Next
	}

	return i
}

//打印数据
func (node *QueueNode) Print() {
	if node == nil {
		return
	}

	node = node.Next
	for node != nil {
		fmt.Print(node.Data, " ")
		node = node.Next
	}

	fmt.Println()
}

//入队 -- 尾插
func (node *QueueNode) Push(data interface{}) {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	newNode := new(QueueNode)
	newNode.Next = nil
	newNode.Data = data

	node.Next = newNode
}

//出队 -- 删除第一个数据结点
func (node *QueueNode) Pop() {
	if node == nil {
		return
	}

	node.Next = node.Next.Next
}

func main() {
	q := new(QueueNode)

	q.Create(1, 2, 3, 4, 5)
	fmt.Println("len :", q.Length())

	q.Print()

	q.Push(6)
	q.Print()
	q.Pop()
	q.Print()
}
