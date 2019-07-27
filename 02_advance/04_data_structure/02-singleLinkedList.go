package main

import (
	"fmt"
	"reflect"
)

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

//创建链表
func (node *LinkNode) Create(data ...interface{}) {
	if node == nil || data == nil || len(data) == 0 {
		return
	}

	for _, v := range data {
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode

		node = node.Next
	}
}

//打印链表 -- 迭代法
func (node *LinkNode) Print1() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
		fmt.Print(node.Data, " ")
	}

	fmt.Println()

}

//打印链表 -- 递归法
//递归因为涉及到大量栈帧保存，所以比迭代占用空间
func (node *LinkNode) Print2() {
	if node == nil {
		return
	}

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}

	node.Next.Print2()

}

//倒序打印链表 -- 递归法
func (node *LinkNode) Print3() {
	if node == nil {
		return
	}

	node.Next.Print3()

	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}

}

//链表的逆序 -- 迭代法
func (node *LinkNode) InverseList1() {
	if node == nil {
		return
	}

	head := node

	var pre *LinkNode //nil
	var next *LinkNode

	node = node.Next
	for node != nil {
		next = node.Next //备份原next指向的结点
		node.Next = pre  //将next指向前一个结点

		pre = node  // 将当前结点赋给pre
		node = next //将node移动到下一个节点
	}

	head.Next = pre //head指向原尾结点
}

//长度
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}

	i := 0
	for node.Next != nil {
		node = node.Next
		i++
	}

	return i
}

//链表的插入--按下标插入
func (node *LinkNode) InsertListByIndex(data interface{}, index int) {
	if node == nil || data == nil || index <= 0 || index > node.Length() {
		return
	}

	pre := node
	for i := 0; i < index; i++ {
		pre = node
		node = node.Next
	}

	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = node

	pre.Next = newNode

}

//链表插入 -- 尾插法
func (node *LinkNode) InsertListByTail(data interface{}) {
	if node == nil || data == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	node.Next = newNode
}

//链表插入 -- 头插法
func (node *LinkNode) InsertListByHead(data interface{}) {
	if node == nil || data == nil {
		return
	}

	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = node.Next

	node.Next = newNode
}

//链表删除 -- 按下标
func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil || index <= 0 || index > node.Length() {
		return
	}

	pre := node
	for i := 0; i < index; i++ {
		pre = node
		node = node.Next
	}

	pre.Next = node.Next

	//GC
	node.Next = nil
	node.Data = nil
	node = nil
}

//链表删除 -- 按数据
func (node *LinkNode) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}

	pre := node
	for node != nil {

		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && reflect.DeepEqual(node.Data, data) {
			pre.Next = node.Next

			//GC
			node.Next = nil
			node.Data = nil
			node = nil

			break
		}

		pre = node
		node = node.Next
	}
}

//链表查找
func (node *LinkNode) Search(data interface{}) int {
	if node == nil || data == nil {
		return -1
	}

	i := 0
	for node.Next != nil {
		i++
		node = node.Next
		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && reflect.DeepEqual(node.Data, data) {
			return i
		}

	}

	return -1
}

//链表销毁
func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Next = nil
	node.Data = nil
	node = nil

}
func main() {
	head := new(LinkNode)

	head.Create(1, 2, 3, 4)

	//head.Print1()
	//head.Print2()
	//head.Print3()

	//fmt.Println("length : ", head.Length())

	head.InverseList1()
	head.Print1()

	//head.InsertListByIndex(3.5, 4)
	//head.InsertListByTail(4.5)
	//head.InsertListByHead(0.1)
	//head.DeleteByIndex(2)
	//head.DeleteByData(3)
	//fmt.Println("index:", head.Search(4))
	head.Destroy()
	fmt.Println("length : ", head.Length())
	head.Print1()
}
