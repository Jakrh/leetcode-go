package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	maxVal := 0
	maxValIdx := 0
	for i, num := range nums {
		if num > maxVal {
			maxVal = num
			maxValIdx = i
		}
	}

	return &TreeNode{
		Val:   maxVal,
		Left:  constructMaximumBinaryTree(nums[:maxValIdx]),
		Right: constructMaximumBinaryTree(nums[maxValIdx+1:]),
	}
}
