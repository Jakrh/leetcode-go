package solution

func missingNumber(nums []int) int {
	sum := (len(nums) + 1) * len(nums) / 2

	for _, n := range nums {
		sum -= n
	}

	return sum
}
