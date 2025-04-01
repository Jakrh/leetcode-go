package solution

func runningSum(nums []int) []int {
	currPrefixSum := 0
	r := make([]int, len(nums))
	for i, n := range nums {
		currPrefixSum += n
		r[i] = currPrefixSum
	}
	return r
}
