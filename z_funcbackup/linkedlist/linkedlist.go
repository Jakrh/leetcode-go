package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertSlice2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	listFirst := &ListNode{
		Val: nums[0],
	}
	listPtr := listFirst
	for i := 1; i < len(nums); i++ {
		listPtr.Next = &ListNode{
			Val: nums[i],
		}
		listPtr = listPtr.Next
	}
	return listFirst
}

func convertList2Slice(list *ListNode) []int {
	s := []int{}
	for list != nil {
		s = append(s, list.Val)
		list = list.Next
	}
	return s
}

func listsEqual(got, want *ListNode) bool {
	if got == nil && want == nil {
		return true
	} else if (got == nil && want != nil) || (got != nil && want == nil) {
		return false
	}

	for got != nil && want != nil {
		if got.Val != want.Val {
			return false
		}
		got = got.Next
		want = want.Next
	}

	return got == want
}
