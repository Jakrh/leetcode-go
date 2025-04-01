package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		values := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			currNode := q[i]
			values = append(values, currNode.Val)

			if currNode.Left != nil {
				q = append(q, currNode.Left)
			}
			if currNode.Right != nil {
				q = append(q, currNode.Right)
			}
		}
		q = q[levelSize:]
		result = append([][]int{values}, result...)
	}

	return result
}
