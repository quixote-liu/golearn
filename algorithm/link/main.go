package main

import (
	"fmt"
	"golearn/algorithm/link/data"
	"golearn/algorithm/link/reverse"
)

func main() {
	ReverseDemo()
}

func ReverseDemo() {
	// reverse
	head := data.GetLinkData()

	// 在没有反转之前
	next := head.Next
	for next != nil {
		fmt.Printf("%d  ", next.Value)
		next = next.Next
	}

	reverse.Reserse(head)

	// 反转之后
	next2 := head.Next
	for next2 != nil {
		fmt.Printf("%d  ", next2.Value)
		next2 = next2.Next
	}
}
