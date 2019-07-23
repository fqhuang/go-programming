package main

import (
	"fmt"
	"reflect"
)

type CircularLinkedList struct {
	Data interface{}
	Next *CircularLinkedList
}

//创建链表
func (node *CircularLinkedList) Create(data ...interface{}) {
	if node == nil || data == nil || len(data) == 0 {
		return
	}
	head := node
	for _, v := range data {
		newNode := new(CircularLinkedList)
		newNode.Next = nil
		newNode.Data = v

		node.Next = newNode
		node = node.Next
	}

	node.Next = head.Next //尾结点指向第一个数据结点
}

//链表长度
func (node *CircularLinkedList) Length() int {
	if node == nil {
		return -1
	}

	start := node.Next
	i := 0

	for {
		i++
		node = node.Next
		if node.Next == start {
			break
		}
	}

	return i
}

//打印链表 -- 正序
func (node *CircularLinkedList) Print1() {
	if node == nil {
		return
	}

	start := node.Next

	for {
		node = node.Next
		fmt.Print(node.Data, " ")
		if node.Next == start {
			break
		}

	}

	fmt.Println()
}

func (node *CircularLinkedList) print(second *CircularLinkedList) {
	if node == nil || node.Next == second {
		return
	}
	node.Next.print(second)

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
}

//打印链表 -- 倒序
func (node *CircularLinkedList) Print2() {
	if node == nil {
		return
	}

	if node.Next != nil {
		defer fmt.Print(node.Next.Data, " ")
	}

	node.Next.Next.print(node.Next.Next)
}

//插入链表
func (node *CircularLinkedList) Insert(data interface{}, index int) {
	if node == nil {
		return
	}

	start := node.Next

	pre := node
	for i := 0; i < index; i++ {
		pre = node
		node = node.Next
	}

	newNode := new(CircularLinkedList)
	newNode.Data = data
	newNode.Next = node

	pre.Next = newNode

	if index == 1 {
		for {
			node = node.Next
			if node.Next == start {
				break
			}
		}

		node.Next = newNode
	}

}

//删除链表 -- 按下标
func (node *CircularLinkedList) DeleteByIndex(index int) {
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	head := node
	start := node.Next
	pre := node
	for i := 0; i < index; i++ {
		pre = node
		node = node.Next
	}

	if index == 1 {
		temp := head
		for {
			temp = temp.Next
			if temp.Next == start {
				temp.Next = node.Next
				break
			}
		}
	}

	pre.Next = node.Next

	node.Next = nil
	node.Data = nil
	node = nil

}

//删除链表 -- 按数据
func (node *CircularLinkedList) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}

	start := node.Next
	pre := node
	node = node.Next
	for node.Next != start {
		if reflect.TypeOf(data) == reflect.TypeOf(node.Data) &&
			reflect.DeepEqual(data, node.Data) {

			if node == start {
				temp := node
				for temp.Next != start {
					temp = temp.Next
				}
				temp.Next = node.Next
			}

			pre.Next = node.Next

			node.Next = nil
			node.Data = nil
			node = nil

			return
		}
		pre = node
		node = node.Next
	}
}

func main() {
	head := new(CircularLinkedList)

	head.Create(1, 2, 3, 4, 5)
	//fmt.Println("length: ", head.Length())
	//head.Print1()

	//head.Insert(0.1, 5)
	//head.DeleteByIndex(1)
	head.DeleteByData(1)
	head.Print1()

}
