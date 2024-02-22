package ztool

import (
	"fmt"
	"math/rand"
	_ "unsafe"
)

type Line struct {
	Value int
	Next  *Line
}

func Rand() {
	randomInt := rand.Intn(100) // 生成一个0到99之间的随机整数
	fmt.Println("Random integer:", randomInt)

	// 再次生成随机整数
	randomInt = rand.Intn(100) // 再次生成一个0到99之间的随机整数
	fmt.Println("Another random integer:", randomInt)
	// 创建一个简单的链表
	//node1 := &Node{data: 1}
	//node2 := &Node{data: 2}
	//node3 := &Node{data: 3}

	//node1.next = node2
	//node2.next = node3
	//
	//fmt.Println("原始链表：")
	//printLinkedList(node1)
	//
	//// 倒序链表
	//reversed := reverseLinkedList(node1)
	//
	//fmt.Println("倒序后的链表：")
	//printLinkedList(reversed)
}

type Node struct {
	data int
	next *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev, next *Node
	current := head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func printLinkedList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

//go:linkname hello gotest/tools.Hello
func hello() {
	println("hello111")
}
