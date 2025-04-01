package solution

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDiameterAndHeight(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lHeight, lDiameter := getDiameterAndHeight(node.Left)
	rHeight, rDiameter := getDiameterAndHeight(node.Right)
	return 1 + maxInt(lHeight, rHeight), maxInt(lHeight+rHeight, maxInt(lDiameter, rDiameter))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := getDiameterAndHeight(root)
	return diameter
}
