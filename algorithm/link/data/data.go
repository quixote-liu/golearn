package data

type LinkHead struct {
	Next *LinkNode
}

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func GetLinkData() *LinkHead {
	node6 := &LinkNode{
		Value: 6,
		Next:  nil,
	}
	node5 := &LinkNode{
		Value: 5,
		Next:  node6,
	}
	node4 := &LinkNode{
		Value: 4,
		Next:  node5,
	}
	node3 := &LinkNode{
		Value: 3,
		Next:  node4,
	}
	node2 := &LinkNode{
		Value: 2,
		Next:  node3,
	}
	node1 := &LinkNode{
		Value: 1,
		Next:  node2,
	}
	h := &LinkHead{
		Next: node1,
	}
	return h
}
