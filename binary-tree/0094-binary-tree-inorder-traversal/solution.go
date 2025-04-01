package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// Runtime: 1ms
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// Return a slice (you can treat a Go slice as an array)
	// that merged by dfs(left), [value], and dfs(right)
	return append(append(
		inorderTraversal1(root.Left),
		root.Val),
		inorderTraversal1(root.Right)...)
}

// DFS: avoiding lot of memory operations for
// the result slice
// Runtime: 0ms
func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0, 100)

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return result
}

// Iteration with a stack
// Runtime: 0ms
func inorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0, 100)
	stack := make([]*TreeNode, 0, 100)

	currNode := root
	for currNode != nil || len(stack) > 0 {
		for currNode != nil {
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		// Pop the last node from stack
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, currNode.Val)
		currNode = currNode.Right
	}

	return result
}
