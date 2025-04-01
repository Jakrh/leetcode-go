package solution

import "testing"

func lookupNodeByValue(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{}
	q = append(q, root)

	var node *TreeNode
	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Val == target {
			return node
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return nil
}

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

func TestLowestCommonAncestor(t *testing.T) {
	type args struct {
		root []any
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
				p:    5,
				q:    1,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
				p:    5,
				q:    4,
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				root: []any{1, 2},
				p:    1,
				q:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := slice2TreeNodes(tt.args.root)
			if got := lowestCommonAncestor(
				root,
				lookupNodeByValue(root, tt.args.p),
				lookupNodeByValue(root, tt.args.q),
			); tt.want != got.Val {
				t.Errorf("func(%v, %v, %v) = %v, want %v",
					tt.args.root,
					tt.args.p,
					tt.args.q,
					got.Val,
					tt.want,
				)
			}
		})
	}
}
