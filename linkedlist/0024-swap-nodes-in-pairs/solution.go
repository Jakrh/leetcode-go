package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	nextNext := head.Next.Next

	next.Next = head
	head.Next = swapPairs(nextNext)

	return next
}
