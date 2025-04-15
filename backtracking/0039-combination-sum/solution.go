package solution

func backtrack(res *[][]int, cand []int, target int, comb []int, sum, startIdx int) {
	if sum > target {
		return
	}

	if sum == target {
		newComb := make([]int, len(comb))
		copy(newComb, comb)
		*res = append(*res, newComb)
		return
	}

	for i := startIdx; i < len(cand); i++ {
		comb = append(comb, cand[i])
		backtrack(res, cand, target, comb, sum+cand[i], i)
		comb = comb[:len(comb)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	combination := make([]int, 0, len(candidates))
	backtrack(&res, candidates, target, combination, 0, 0)
	return res
}
