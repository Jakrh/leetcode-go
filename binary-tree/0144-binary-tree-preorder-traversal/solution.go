package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// Runtime: 0ms
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// Return a slice (you can treat a Go slice as an array)
	// that merged by [value], dfs(left), and dfs(right)
	return append(append(
		[]int{root.Val},
		preorderTraversal1(root.Left)...),
		preorderTraversal1(root.Right)...)
}

// DFS: avoiding lot of memory operations for
// the result slice
// Runtime: 0ms
func preorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0, 100)

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return result
}

// Iteration with a stack
// Runtime: 0ms
func preorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0, 100)
	stack := make([]*TreeNode, 0, 100)

	currNode := root
	for currNode != nil || len(stack) > 0 {
		for currNode != nil {
			result = append(result, currNode.Val)
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		// Pop the last node from stack
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currNode = currNode.Right
	}

	return result
}
