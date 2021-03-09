package main

import "fmt"

// 单向链表
func main() {
	// 创建链表
	SplitLine("创建链表")
	CreateLT()

	// 插入节点
	// 头插法
	SplitLine("头插法")
	HeadInsert()
	// 尾插法
	SplitLine("尾插法")
	TailInsert()
}

type Node struct {
	Data int
	Next *Node
}

func ShowNode(n *Node) {
	for n != nil {
		fmt.Printf("Node:%p,value:%d\n", n, n.Data)
		n = n.Next
	}
}

func CreateLT() {
	// 创建头结点
	var head = new(Node)
	head.Data = 1
	// 添加节点
	var node1 = new(Node)
	node1.Data = 2

	head.Next = node1

	var node2 = new(Node)
	node2.Data = 3

	node1.Next = node2
	ShowNode(head)

}

func HeadInsert() {

	// 头插法

	var head = new(Node)
	head.Data = 1
	var tail *Node

	tail = head

	for i := 0; i < 10; i++ {
		node := Node{Data: i}
		node.Next = tail
		tail = &node
	}

	//ShowNode(head)
	ShowNode(tail)
}

func TailInsert() {
	var head = new(Node)
	head.Data = 1
	var tail *Node
	tail = head
	for i := 0; i < 10; i++ {
		var node = Node{Data: i}
		(*tail).Next = &node
		tail = &node
	}
	ShowNode(head)
}

func SplitLine(info string) {
	fmt.Printf("------------%s-------------\n", info)
}
