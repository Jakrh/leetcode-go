package main

import (
	"testing"
)

func BenchmarkMergeTwoLists(b *testing.B) {
	list1 := convertSlice2List([]int{1, 2, 4})
	list2 := convertSlice2List([]int{1, 3, 4})
	for i := 0; i < b.N; i++ {
		mergeTwoLists(list1, list2)
	}
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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "1",
			args: [][]int{{1, 2, 4}, {1, 3, 4}},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := convertSlice2List(tt.args[0])
			list2 := convertSlice2List(tt.args[1])
			want := convertSlice2List(tt.want)
			if got := mergeTwoLists(list1, list2); !listsEqual(got, want) {
				t.Errorf("mergeTwoLists() = %v, want %v", convertList2Slice(got), tt.want)
			}
		})
	}
}
