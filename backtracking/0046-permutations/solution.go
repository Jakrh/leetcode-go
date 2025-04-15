package solution

func backtrack(res *[][]int, nums []int, premutation []int, used []bool) {
	if len(premutation) == len(nums) {
		newPremutation := make([]int, len(premutation))
		copy(newPremutation, premutation)
		*res = append(*res, newPremutation)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			premutation = append(premutation, nums[i])
			backtrack(res, nums, premutation, used)
			used[i] = false
			premutation = premutation[:len(premutation)-1]
		}
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	premutation := make([]int, 0, len(nums))
	backtrack(&res, nums, premutation, used)
	return res
}
