package main

import (
	"testing"
)

// for this special question to set the cycle
func convertSlice2List(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	listFirst := &ListNode{
		Val: nums[0],
	}

	// set the cycle if pos is 0
	var cycledNode *ListNode
	if pos == 0 {
		cycledNode = listFirst
	}

	listPtr := listFirst
	for i := 1; i < len(nums); i++ {
		listPtr.Next = &ListNode{
			Val: nums[i],
		}
		listPtr = listPtr.Next

		// set the cycle
		if i == pos {
			cycledNode = listPtr
		}
		if i == len(nums)-1 && cycledNode != nil {
			listPtr.Next = cycledNode
		}
	}

	return listFirst
}

func TestHasCycle(t *testing.T) {
	type args struct {
		list []int
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				list: []int{3, 2, 0, -4},
				pos:  1,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				list: []int{1, 2},
				pos:  0,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				list: []int{1},
				pos:  -1,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				list: []int{},
				pos:  -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := convertSlice2List(tt.args.list, tt.args.pos)
			if got := hasCycle(l); got != tt.want {
				t.Errorf("func(%v) = %v, want %v", tt.args.list, got, tt.want)
			}
		})
	}
}
