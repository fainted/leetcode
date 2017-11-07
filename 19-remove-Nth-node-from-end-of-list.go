package main

// 快慢指针
// [] = logList(removeNthFromEnd(nil, 0))
// [] = logList(removeNthFromEnd(&ListNode{1, nil}, 1))
// [1, 2, 3, 4] = logList(removeNthFromEnd(listFromSlice([]int{1, 2, 3, 4, 5}), 2))
// [2, 3, 4, 5] = logList(removeNthFromEnd(listFromSlice([]int{1, 2, 3, 4, 5}), 5))
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{0, head}
	fast, slow := &dummy, &dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	if slow != nil && slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
