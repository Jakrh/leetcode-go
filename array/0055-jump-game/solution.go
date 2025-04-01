package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canJump(nums []int) bool {
	maxIdx := 0
	lastIdx := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > maxIdx {
			maxIdx = nums[i] + i
		}
		if maxIdx == i && i != lastIdx {
			return false
		}
	}

	return true
}
