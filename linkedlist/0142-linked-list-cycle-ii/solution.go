package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	cycleFound := false

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			cycleFound = true
			break
		}
	}
	if !cycleFound {
		return nil
	}

	slow2 := head
	for slow != slow2 {
		slow = slow.Next
		slow2 = slow2.Next
	}
	return slow
}
