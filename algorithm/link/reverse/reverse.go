package reverse

import "golearn/algorithm/link/data"

// 使用原始的方法达到原地反转功能
func Reserse(head *data.LinkHead) {
	if head == nil || head.Next == nil {
		return
	}

	var pre *data.LinkNode  // 前驱节点
	var back *data.LinkNode // 后驱节点
	next := head.Next
	for next != nil {
		back = next.Next
		next.Next = pre
		pre = next
		next = back
	}
	head.Next = pre
}
