package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := invertTree1(root.Left)
	r := invertTree1(root.Right)

	root.Left = r
	root.Right = l

	return root
}

// BFS
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 0, 101)
	var currNode *TreeNode

	q = append(q, root)
	for len(q) > 0 {
		currNode = q[0]
		q = q[1:]
		currNode.Left, currNode.Right = currNode.Right, currNode.Left
		if currNode.Left != nil {
			q = append(q, currNode.Left)
		}
		if currNode.Right != nil {
			q = append(q, currNode.Right)
		}
	}

	return root
}
