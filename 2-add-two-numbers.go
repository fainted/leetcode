package leetcode

func digit(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func next(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var head ListNode
	tail := &head
	carry := int(0)

	for c1, c2 := l1, l2; c1 != nil || c2 != nil; c1, c2 = next(c1), next(c2) {
		sum := digit(c1) + digit(c2) + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		tail.Next = &ListNode{sum, nil}
		tail = tail.Next
	}

	if carry == 1 {
		tail.Next = &ListNode{1, nil}
	}

	return head.Next
}
