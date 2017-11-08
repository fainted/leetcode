package main

func reverseList(head *ListNode) *ListNode {
	var dummy ListNode
	for cur := head; cur != nil; {
		p := cur
		p.Next, cur = dummy.Next, cur.Next
		dummy.Next = p
	}

	return dummy.Next
}
