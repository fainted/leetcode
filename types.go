package main

import "log"

// ListNode linked list node definition
type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{nums[0], nil}
	for i, p := 1, head; i < len(nums); i++ {
		p.Next = &ListNode{nums[i], nil}
		p = p.Next
	}

	return head
}

func logList(head *ListNode) {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	log.Print(nums)
}
