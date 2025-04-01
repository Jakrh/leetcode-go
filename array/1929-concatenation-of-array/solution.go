package solution

func getConcatenation1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)

	idx := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			ans[idx] = nums[j]
			idx++
		}
	}

	return ans
}

func getConcatenation2(nums []int) []int {
	n := len(nums)
	ans := make([]int, 0, n*2)

	for i := 0; i < 2; i++ {
		ans = append(ans, nums...)
	}

	return ans
}

func getConcatenation3(nums []int) []int {
	return append(nums, nums...)
}
