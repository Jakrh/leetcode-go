package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func getMaxDepth(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	return 1 + maxInt(getMaxDepth(node.Left, max), getMaxDepth(node.Right, max))
}

func maxDepth(root *TreeNode) int {
	return getMaxDepth(root, 0)
}
