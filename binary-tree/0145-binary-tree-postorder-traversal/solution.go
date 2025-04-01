package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// Runtime: 1ms
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// Return a slice (you can treat a Go slice as an array)
	// that merged by dfs(left), and dfs(right), [value]
	return append(append(
		postorderTraversal1(root.Left),
		postorderTraversal1(root.Right)...),
		root.Val)
}

// DFS: avoiding lot of memory operations for
// the result slice
// Runtime: 0ms
func postorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0, 100)

	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		result = append(result, root.Val)
	}

	postorder(root)
	return result
}

// Iteration with a stack
// Runtime: 0ms
func postorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0, 100)
	stack := make([]*TreeNode, 0, 100)

	deepestNode := root
	var prevNode *TreeNode
	for deepestNode != nil || len(stack) > 0 {
		for deepestNode != nil {
			stack = append(stack, deepestNode)
			deepestNode = deepestNode.Left
		}

		currNode := stack[len(stack)-1]
		if currNode.Right != nil && currNode.Right != prevNode {
			deepestNode = currNode.Right
		} else {
			// Pop the deepest node from stack
			prevNode = currNode
			stack = stack[:len(stack)-1]

			result = append(result, prevNode.Val)
		}
	}

	return result
}
