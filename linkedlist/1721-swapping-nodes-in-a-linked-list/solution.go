package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	left := head
	for i := 1; i < k; i++ {
		left = left.Next
	}

	right := head
	currNode := left
	for currNode.Next != nil {
		currNode = currNode.Next
		right = right.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return head
}
