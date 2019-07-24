package main

import (
	"fmt"
	"reflect"
)

type BinTreeNode struct {
	Data  interface{}
	LNode *BinTreeNode
	RNode *BinTreeNode
}

//新建二叉树
func (node *BinTreeNode) Create() {

	node1 := BinTreeNode{Data: 1, LNode: nil, RNode: nil}
	node2 := BinTreeNode{Data: 2, LNode: nil, RNode: nil}
	node3 := BinTreeNode{Data: 3, LNode: nil, RNode: nil}
	node4 := BinTreeNode{Data: 4, LNode: nil, RNode: nil}
	node5 := BinTreeNode{Data: 5, LNode: nil, RNode: nil}
	//node6 := BinTreeNode{Data: 6, LNode: nil, RNode: nil}
	//node7 := BinTreeNode{Data: 7, LNode: nil, RNode: nil}

	node.Data = 0
	node.LNode = &node1
	node.RNode = &node2
	node1.LNode = &node3
	node1.RNode = &node4
	node2.LNode = &node5
	//node2.RNode = &node6
	//node3.LNode = &node7

}

//打印二叉树 -- 先序遍历:DLR
func (node *BinTreeNode) PreOrder() {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")

	node.LNode.PreOrder()

	node.RNode.PreOrder()
}

//打印二叉树 -- 中序遍历:LDR
func (node *BinTreeNode) MidOrder() {
	if node == nil {
		return
	}

	node.LNode.MidOrder()

	fmt.Print(node.Data, " ")

	node.RNode.MidOrder()
}

//打印二叉树 -- 后序遍历:LRD
func (node *BinTreeNode) PostOrder() {
	if node == nil {
		return
	}

	node.LNode.PostOrder()

	node.RNode.PostOrder()

	fmt.Print(node.Data, " ")
}

//获取二叉树的深度（高度）
func (node *BinTreeNode) Height() int {
	if node == nil {
		return 0 //不能返回-1
	}

	lh := node.LNode.Height()
	rh := node.RNode.Height()

	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}

}

//获取叶子结点个数 -- 指针传递
func (node *BinTreeNode) LeafNum(num *int) {
	if node == nil {
		return
	}

	if node.LNode == nil && node.RNode == nil {
		*num++
	}

	node.LNode.LeafNum(num)
	node.RNode.LeafNum(num)

}

//获取叶子结点个数 -- 全局变量
var num2 = 0

func (node *BinTreeNode) LeafNum2() {
	if node == nil {
		return
	}

	if node.LNode == nil && node.RNode == nil {
		num2++
	}

	node.LNode.LeafNum2()
	node.RNode.LeafNum2()
}

//查找二叉树
var bl = false

func (node *BinTreeNode) Search(data interface{}) {
	if node == nil || data == nil {
		return
	}

	if reflect.TypeOf(data) == reflect.TypeOf(node.Data) &&
		reflect.DeepEqual(data, node.Data) {
		bl = true
		fmt.Println("find data:", node.Data)
		return
	}

	node.LNode.Search(data)
	node.RNode.Search(data)
}

//翻转二叉树
func (node *BinTreeNode) Reverse() {
	if node == nil {
		return
	}

	node.LNode, node.RNode = node.RNode, node.LNode

	node.LNode.Reverse()
	node.RNode.Reverse()
}

//拷贝二叉树
func (node *BinTreeNode) Copy() *BinTreeNode {
	if node == nil {
		return nil
	}

	lNode := node.LNode.Copy()
	rNode := node.RNode.Copy()

	newNode := new(BinTreeNode)
	newNode.Data = node.Data
	newNode.LNode = lNode
	newNode.RNode = rNode

	return newNode
}

//销毁二叉树
func (node *BinTreeNode) Destroy() {
	if node == nil {
		return
	}

	node.LNode.Destroy()
	node.LNode = nil

	node.RNode.Destroy()
	node.RNode = nil
	node.Data = nil

}
func main() {
	tree := new(BinTreeNode)
	tree.Create()

	fmt.Println(tree)

	tree.PreOrder()
	fmt.Println()
	tree.MidOrder()
	fmt.Println()
	tree.PostOrder()
	fmt.Println()

	fmt.Println("height:", tree.Height())

	num := new(int)
	tree.LeafNum(num)
	fmt.Println("leaf num :", *num)
	tree.LeafNum2()
	fmt.Println("leaf num2:", num2)

	tree.Search(4)

	tree.Reverse()
	tree.PreOrder()
	fmt.Println()

	newTree := tree.Copy()
	newTree.PreOrder()
	fmt.Println()

	tree.Destroy()
	tree.PreOrder()
	//newTree.PreOrder()
}
