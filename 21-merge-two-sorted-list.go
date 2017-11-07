package main

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy

	for p1, p2 := l1, l2; p1 != nil || p2 != nil; {
		var cur **ListNode
		if (p1 != nil && p2 != nil && p1.Val <= p2.Val) || (p1 != nil && p2 == nil) {
			cur = &p1
		} else {
			cur = &p2
		}

		tail.Next = &ListNode{(*cur).Val, nil}
		tail, *cur = tail.Next, (*cur).Next
	}

	return dummy.Next
}