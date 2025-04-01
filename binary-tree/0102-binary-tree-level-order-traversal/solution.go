package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)

	var node *TreeNode
	var currSize int
	for len(q) > 0 {
		currSize = len(q)
		currLevel := []int{}
		for i := 0; i < currSize; i++ {
			// Pub a node from queue
			node = q[0]
			q = q[1:]

			currLevel = append(currLevel, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		result = append(result, currLevel)
	}

	return result
}
