package goLeet

// Definition for singly-linked list.
type ListNode struct {
     Val int
    Next *ListNode
}

func getValue (l* ListNode, placeValue uint64) uint64 {
	if l == nil {
		return 0
	}
	return placeValue * uint64(l.Val) + getValue(l.Next, placeValue * 10)
}

func listNodeToUInt(l *ListNode) uint64 {
	val := getValue(l, 1)
	return val
}

func uintToListNode(i uint64) *ListNode {
	l := &ListNode{Val: 0, Next: nil}
	root := l
	for {
		l.Val = int(i % 10)
		i = i / 10
		if i == 0 {
			break
		}
		l2 := &ListNode{Val: 0, Next: nil}
		l.Next = l2
		l = l2
	}
	return root
}

func addTwoNumbersUint(l1 *ListNode, l2 *ListNode) *ListNode {
	i1 := listNodeToUInt(l1)
	i2 := listNodeToUInt(l2)
	sum := i1 + i2
	return uintToListNode(sum)
}