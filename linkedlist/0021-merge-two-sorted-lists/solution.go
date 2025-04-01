package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode

	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head = list1
			curr = head
			list1 = list1.Next
		} else {
			head = list2
			curr = head
			list2 = list2.Next
		}
	} else if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	} else {
		curr.Next = nil
	}

	return head
}
