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

func TestRemoveNodes(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				list: []int{5, 2, 13, 3, 8},
			},
			want: []int{13, 8},
		},
		{
			name: "2",
			args: args{
				list: []int{1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := convertSlice2List(tt.args.list)
			want := convertSlice2List(tt.want)
			if got := removeNodes(l); !listsEqual(got, want) {
				t.Errorf("func(%v) = %v, want %v", tt.args.list, convertList2Slice(got), tt.want)
			}
		})
	}
}
