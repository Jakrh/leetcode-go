package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func intSlice2AnySlice(nums []int) []any {
	iface := make([]any, len(nums))
	for i, n := range nums {
		iface[i] = n
	}
	return iface
}

func int2dSlice2AnySlice(nums [][]int) [][]any {
	iface := make([][]any, len(nums))
	for i, s := range nums {
		iface[i] = intSlice2AnySlice(s)
	}
	return iface
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

func compare2dSlices(a, b [][]any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !compareSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

func treesEqual(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if (tree1 == nil && tree2 != nil) || (tree1 != nil && tree2 == nil) {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return treesEqual(tree1.Left, tree2.Left) && treesEqual(tree1.Right, tree2.Right)
}
