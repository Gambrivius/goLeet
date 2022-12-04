package goLeet

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lSum := &ListNode{Val: 0, Next: nil}
	root := lSum
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			lSum.Val += l1.Val
		}
		if l2 != nil {
			lSum.Val += l2.Val
		}
		lSum.Val += carry
		carry = 0
		if lSum.Val >= 10 {
			carry = 1
			lSum.Val = lSum.Val - 10
		}
		if l1 != nil {
            l1 = l1.Next
        }
		if l2 != nil {
            l2 = l2.Next
        }
        if l1 != nil || l2 != nil {
		    lSum.Next = &ListNode{Val: 0, Next: nil}
		    lSum = lSum.Next
        }
	}
	if carry == 1 {
		lSum.Next = &ListNode{Val: 1, Next: nil}
	}
	return root
}