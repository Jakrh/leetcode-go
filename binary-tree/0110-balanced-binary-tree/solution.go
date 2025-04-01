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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getHeight(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	lHeight, lBalanced := getHeight(node.Left)
	rHeight, rBalanced := getHeight(node.Right)
	return 1 + maxInt(lHeight, rHeight), (absInt(lHeight-rHeight) <= 1) && (lBalanced && rBalanced)
}

func isBalanced(root *TreeNode) bool {
	_, balanced := getHeight(root)
	return balanced
}
