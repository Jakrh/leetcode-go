package solution

import (
	"reflect"
	"testing"
)

func treeNodes2Slice(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	s := []any{}
	q := []*TreeNode{}
	var currNode *TreeNode

	q = append(q, root)
	for len(q) > 0 {
		currNode = q[0]
		q = q[1:]

		if currNode != nil {
			s = append(s, currNode.Val)
			q = append(q, currNode.Left)
			q = append(q, currNode.Right)
		} else {
			s = append(s, nil)
		}
	}

	// Trim all tailing nils
	trimPos := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == nil {
			trimPos--
		} else {
			break
		}
	}
	s = s[:trimPos]

	return s
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			want: []any{6, 3, 5, nil, 2, 0, nil, nil, 1},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []any{3, nil, 2, nil, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(tt.want, treeNodes2Slice(got)) {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, treeNodes2Slice(got), tt.want)
			}
		})
	}
}
