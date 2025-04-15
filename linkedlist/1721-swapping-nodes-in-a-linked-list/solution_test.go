package main

import (
	"testing"
)

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

func TestSwapNodes(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "2",
			args: args{
				list: []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5},
				k:    5,
			},
			want: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := convertSlice2List(tt.args.list)
			want := convertSlice2List(tt.want)
			if got := swapNodes(l, tt.args.k); !listsEqual(got, want) {
				t.Errorf("func(%v) = %v, want %v", tt.args.list, convertList2Slice(got), tt.want)
			}
		})
	}
}
