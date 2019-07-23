package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
) // 导入C代码

// 定义切片类型
type Slice struct {
	Data unsafe.Pointer // Go语言中的 万能指针类型. void * C语言中的万能指针. 没有具体数据类型,go不能进行运算.
	Len  int            // 数据元素个数
	Cap  int            // 可扩展的有效容量
}

const TAG = 8

// 绑定创建 切片 方法	--- Create(5, 5, 1,2,3,4,5)
func (s *Slice) Create(l int, c int, Data ...int) {
	// 容错处理
	if s == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}
	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}

	// 申请内存空间 -- 单位 : 字节  --- 返回 void *类型 不能参与运算  s.Data是一个地址值.(不能运算)
	s.Data = C.malloc(C.size_t(c) * 8)

	// 初始化 长度和容量
	s.Len = l
	s.Cap = c
	// 将 s.Data 转换成 可以计算的 数值.	将 0x420800 地址 变成数值 0x420800
	p := uintptr(s.Data)

	// 根据 Data 集合,遍历存入申请的内存中
	for _, v := range Data {
		// 将 数值 p,转换回地址值. 具体数据类型, 解引用 获取内存
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}
	// 释放内存
	// C.free(s.Data)
}

// 打印切片方法
func (s *Slice) Print() {
	if s == nil {
		return
	}
	// 将 地址转换为可以计算的 数值
	p := uintptr(s.Data)

	// 按 len 循环 打印切片元素
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += TAG
	}
	fmt.Println()
}

// 追加元素 Append(数据)		--- Append(666, 777, 888)
func (s *Slice) Append(Data ...int) {
	if s == nil {
		return
	}
	// 判断是否需要扩容
	for len(Data)+s.Len > s.Cap {
		// 拓展容量为原来的 2 倍, 存储新内存空间地址
		s.Data = C.realloc(s.Data, C.size_t(s.Cap)*2*8)
		s.Cap *= 2
	}
	// 将s.Data 转换程可以运算的 数值
	p := uintptr(s.Data)

	/*	for i:=0; i<s.Len; i++ {
		p += TAG
	}*/
	// 偏移 p ,到 结尾处
	p += uintptr(s.Len) * 8

	// 按 data 循环,依次取出数据, 存入内存中
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}

	// 修改 len
	s.Len += len(Data)
}

//根据切片下标取元素
func (s *Slice) GetData(index int) int {
	if s == nil {
		return -1
	}
	if index < 0 || index >= s.Len {
		return -1
	}
	// 将 万能指针转换为可以计算的 数据
	p := uintptr(s.Data)

	// 偏移p,到 index 指代的元素位置
	p += uintptr(index) * TAG

	// 取数据值,返回
	return *(*int)(unsafe.Pointer(p))
}

// 已知元素,返回下标值
func (s *Slice) SearchData(Data int) int {
	if s == nil {
		return -1
	}
	// 将万能指针,转换为可以计算的数据值
	p := uintptr(s.Data)

	for i := 0; i < s.Len; i++ {
		if *(*int)(unsafe.Pointer(p)) == Data {
			return i
		}
		p += TAG
	}
	return -1
}

// 根据下标,删除切片元素
func (s *Slice) Delete(index int) {
	if s == nil {
		return
	}
	if index < 0 || index >= s.Len {
		return
	}
	// 将万能指针,转换为可以计算的数据值
	p := uintptr(s.Data)

	// 将p 偏移到 index 位置
	p += uintptr(index) * TAG

	// 定义 变量,记录 p 指代的元素的 下一个元素
	aftp := p

	// 循环从 index 到 s.Len 依次完成 后一个元素给前一个元素赋值
	for i := index; i < s.Len; i++ {
		aftp += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(aftp))
		p += TAG
	}
	// 修改 s.Len 去除一个元素
	s.Len -= 1
}

// 根据index ,向 切片插入数据
func (s *Slice) Insert(Data int, index int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.Len {
		return
	}
	// 如果 插入的index 在切片结尾
	if index == s.Len {
		s.Append(Data)
		return
	}

	// 判断是否超出cap

	// 如果插入位置 在 中间
	p := uintptr(s.Data)

	// 将p偏移到 index 位置
	p += uintptr(index) * TAG

	// 获取插入元素完成后的,最后一个元素位置.
	temp := uintptr(s.Data) + uintptr(s.Len)*TAG

	// 循环将 index 之后的元素依次后移(前一个元素,给后一个元素赋值)
	for i := s.Len; i > index; i-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}

	// 循环结束后,将 p 对应的内存,写入 参数 Data
	*(*int)(unsafe.Pointer(p)) = Data

	// 修改 s.Len
	s.Len++
}

// 销毁 切片
func (s *Slice) Destroy() {
	if s == nil || s.Data == nil {
		return
	}
	C.free(s.Data)
	s.Data = nil // 驱使go GC工作
	s.Len = 0
	s.Cap = 0
	s = nil

	runtime.GC() // 手动调用 GC
}

func main() {
	slice := new(Slice)
	// 创建切片
	slice.Create(5, 5, 1, 2, 3, 4, 5)
	//fmt.Println(slice)
	slice.Print()

	// 追加元素
	slice.Append(6, 7, 8, 9, 10)

	slice.Print()
	fmt.Println("长度:", slice.Len)
	fmt.Println("容量:", slice.Cap)

	// 获取元素
	ret := slice.GetData(1)
	fmt.Println("ret=", ret)

	// 给定元素,获取下标值
	idx := slice.SearchData(125)
	fmt.Println("下标:", idx)

	// 删除元素
	slice.Delete(15)
	slice.Print()

	// 插入原素
	slice.Insert(666, 9)
	slice.Print()
	fmt.Println("长度:", slice.Len)
	fmt.Println("容量:", slice.Cap)

	fmt.Println("销毁切片:")
	slice.Destroy()
	slice.Print()
}
