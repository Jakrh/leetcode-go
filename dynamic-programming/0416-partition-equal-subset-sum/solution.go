package solution

func sumInt(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func canPartition(nums []int) bool {
	sum := sumInt(nums)
	if sum&1 == 1 {
		return false
	}
	halfSum := sum / 2
	dp := make([]bool, halfSum+1)

	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := halfSum; j >= i; j-- {
			if nums[i-1] <= j && dp[j-nums[i-1]] {
				dp[j] = true
			}
		}
	}

	return dp[halfSum]
}
