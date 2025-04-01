package solution

import "testing"

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

func TestBuildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []any{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name: "2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: []any{-1},
		},
		{
			name: "3",
			args: args{
				preorder: []int{3, 9, 4, 20, 15, 7},
				inorder:  []int{4, 9, 3, 15, 20, 7},
			},
			want: []any{3, 9, 20, 4, nil, 15, 7},
		},
		{
			name: "4",
			args: args{
				preorder: []int{3, 9, 4, 8, 20, 15, 7},
				inorder:  []int{4, 9, 8, 3, 15, 20, 7},
			},
			want: []any{3, 9, 20, 4, 8, 15, 7},
		},
		{
			name: "5",
			args: args{
				preorder: []int{3, 9, 4, 1, 8, 20, 15, 7},
				inorder:  []int{1, 4, 9, 8, 3, 15, 20, 7},
			},
			want: []any{3, 9, 20, 4, 8, 15, 7, 1},
		},
		{
			name: "6",
			args: args{
				preorder: []int{},
				inorder:  []int{},
			},
			want: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !compareSlices(tt.want, treeNodes2Slice(got)) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.preorder, tt.args.inorder, treeNodes2Slice(got), tt.want)
			}
		})
	}
}
