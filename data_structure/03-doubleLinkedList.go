package main

import (
	"fmt"
	"reflect"
)

type DoubleLinkedList struct {
	Data interface{}
	Pre  *DoubleLinkedList
	Next *DoubleLinkedList
}

//创建链表
func (node *DoubleLinkedList) Create(data ...interface{}) {
	if node == nil || data == nil || len(data) == 0 {
		return
	}

	for _, v := range data {

		newNode := new(DoubleLinkedList)
		newNode.Data = v
		newNode.Next = nil
		newNode.Pre = node

		node.Next = newNode
		node = newNode

	}
}

//链表长度
func (node *DoubleLinkedList) Length() int {
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

//打印链表 -- 迭代法
func (node *DoubleLinkedList) Print1() {
	if node == nil {
		return
	}

	node = node.Next
	for node != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Next
	}

	fmt.Println()
}

//打印链表 -- 递归法
func (node *DoubleLinkedList) Print2() {
	if node == nil {
		return
	}

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}

	node.Next.Print2()

}

//倒叙打印链表 -- 迭代法
func (node *DoubleLinkedList) Print3() {
	if node == nil {
		return
	}

	head := node

	for node.Next != nil {
		node = node.Next
	}

	for node != head {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Pre
	}

	fmt.Println()
}

//倒叙打印链表 -- 递归法
func (node *DoubleLinkedList) Print4() {
	if node == nil {
		return
	}

	node.Next.Print4()

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
}

//链表逆序 -- 迭代法
func (node *DoubleLinkedList) Inverse() {
	if node == nil {
		return
	}

	head := node
	var pre *DoubleLinkedList //nil
	var next *DoubleLinkedList

	node = node.Next //跳过头结点
	for node != nil {
		next = node.Next
		node.Next = pre
		node.Pre = next

		pre = node
		node = next

	}

	pre.Pre = head
	head.Next = pre
}

//链表插入
func (node *DoubleLinkedList) Insert(data interface{}, index int) {
	if node == nil || data == nil || index <= 0 || index > node.Length() {
		return
	}

	for i := 0; i < index; i++ {
		node = node.Next
	}

	newNode := new(DoubleLinkedList)
	newNode.Data = data
	newNode.Pre = node.Pre
	newNode.Next = node

	node.Pre.Next = newNode
	node.Pre = newNode

}

//链表插入--尾插法
func (node *DoubleLinkedList) Insert2(data interface{}) {
	if node == nil || data == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	newNode := new(DoubleLinkedList)
	newNode.Data = data
	newNode.Next = nil
	newNode.Pre = node

	node.Next = newNode
}

//链表删除 -- 根据下标删除
func (node *DoubleLinkedList) DeleteByIndex(index int) {
	if node == nil || index <= 0 || index > node.Length() {
		return
	}

	pre := node
	for i := 0; i < index; i++ {
		pre = node
		node = node.Next
	}

	pre.Next = node.Next
	if node.Next != nil { //删除尾结点时next为nil
		node.Next.Pre = pre
	}

	//GC
	node.Pre = nil
	node.Next = nil
	node.Data = nil
	node = nil
}

//链表删除 -- 按数据删除
func (node *DoubleLinkedList) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}

	for node.Next != nil {
		pre := node
		node = node.Next
		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && reflect.DeepEqual(data, node.Data) {
			pre.Next = node.Next
			if node.Next != nil {
				node.Next.Pre = pre
			}

			//GC
			node.Pre = nil
			node.Next = nil
			node.Data = nil
			node = nil

			return
		}
	}
}

//链表查找
func (node *DoubleLinkedList) Search(data interface{}) int {
	if node == nil || data == nil {
		return -1
	}

	i := 0
	for node.Next != nil {
		i++
		node = node.Next
		if reflect.TypeOf(data) == reflect.TypeOf(node.Data) || reflect.DeepEqual(data, node.Data) {
			return i
		}
	}

	return -1
}

//链表销毁
func (node *DoubleLinkedList) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Pre = nil
	node.Next = nil
	node.Data = nil
}

func main() {
	head := new(DoubleLinkedList)

	head.Create(1, 2, 3, 4, 5)
	fmt.Println("length: ", head.Length())
	//head.Print1()
	//head.Print2()
	//head.Print3()
	//head.Print4()

	//head.Inverse()
	//head.Print1()

	//head.Insert(6, 5)
	//head.DeleteByIndex(5)
	//head.Insert2(6)
	//head.DeleteByData(3)
	//fmt.Println("find index:", head.Search(3))
	head.Destroy()
	head.Print1()
}
