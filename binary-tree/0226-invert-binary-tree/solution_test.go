package solution

import (
	"testing"
)

func slice2TreeNodes(s []any) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	var root *TreeNode
	var val int
	var ok bool
	if val, ok = s[0].(int); ok {
		root = &TreeNode{
			Val: val,
		}
	} else {
		return nil
	}

	q := []*TreeNode{}
	q = append(q, root)
	var currNode *TreeNode
	i, j := 1, 2
	for i < len(s) || j < len(s) {
		if len(q) == 0 {
			break
		}
		currNode = q[0]
		q = q[1:]
		if i < len(s) {
			if val, ok = s[i].(int); ok {
				currNode.Left = &TreeNode{
					Val: val,
				}
				q = append(q, currNode.Left)
			}
		}
		if j < len(s) {
			if val, ok = s[j].(int); ok {
				currNode.Right = &TreeNode{
					Val: val,
				}
				q = append(q, currNode.Right)
			}
		}
		i += 2
		j += 2
	}

	return root
}

func treeNodes2Slice(root *TreeNode) []interface{} {
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

func compareSlices(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type args struct {
	root []any
}

var tests = []struct {
	name string
	args args
	want []any
}{
	{
		name: "1",
		args: args{
			root: []any{4, 2, 7, 1, 3, 6, 9},
		},
		want: []any{4, 7, 2, 9, 6, 3, 1},
	},
	{
		name: "2",
		args: args{
			root: []any{4, 2, 7, 1, 3, 6, 9},
		},
		want: []any{4, 7, 2, 9, 6, 3, 1},
	},
	{
		name: "3",
		args: args{
			root: []any{},
		},
		want: []any{},
	},
	{
		name: "4",
		args: args{
			root: []any{4, 2, 7, 1, 3, nil, 9},
		},
		want: []any{4, 7, 2, 9, nil, 3, 1},
	},
}

func TestInvertTree1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree1(slice2TreeNodes(tt.args.root)); !compareSlices(tt.want, treeNodes2Slice(got)) {
				t.Errorf("func(%v) = %v, want %v", tt.args.root, treeNodes2Slice(got), tt.want)
			}
		})
	}
}

func TestInvertTree2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree2(slice2TreeNodes(tt.args.root)); !compareSlices(tt.want, treeNodes2Slice(got)) {
				t.Errorf("func(%v) = %v, want %v", tt.args.root, treeNodes2Slice(got), tt.want)
			}
		})
	}
}
