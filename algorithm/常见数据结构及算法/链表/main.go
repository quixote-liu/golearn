package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	node := new(LinkNode)
	node.Data = 10

	node1 := new(LinkNode)
	node1.Data = 20
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 30
	node1.NextNode = node2

	// 遍历链表
	nowNode := node
	for nowNode != nil {
		fmt.Println("data:", nowNode.Data)
		nowNode = nowNode.NextNode
	}
}
