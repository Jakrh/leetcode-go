package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := &ListNode{}
	currNode := head

	for l1 != nil || l2 != nil || carry != 0 {
		var sum int

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		val := sum % 10
		carry = sum / 10

		currNode.Next = &ListNode{
			Val: val,
		}
		currNode = currNode.Next
	}

	return head.Next
}
