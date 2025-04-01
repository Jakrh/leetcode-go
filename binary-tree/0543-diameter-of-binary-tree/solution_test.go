package solution

import "testing"

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

func TestDiameterOfBinaryTree(t *testing.T) {
	type args struct {
		tree []any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				tree: []any{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				tree: []any{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(slice2TreeNodes(tt.args.tree)); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.tree, got, tt.want)
			}
		})
	}
}
