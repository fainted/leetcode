package main

func mergeKLists(lists []*ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy

	for {
		var cur **ListNode
		for i := 0; i < len(lists); i++ {
			temp := &lists[i]
			if *temp != nil && (cur == nil || (*temp).Val < (*cur).Val) {
				cur = temp
			}
		}

		if cur == nil {
			break
		}
		tail.Next = &ListNode{(*cur).Val, nil}
		tail, *cur = tail.Next, (*cur).Next
	}

	return dummy.Next
}
