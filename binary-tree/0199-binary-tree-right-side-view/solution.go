package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	q := []*TreeNode{}

	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			node := q[0]
			q = q[1:]

			if size == 1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			size--
		}
	}

	return res
}
