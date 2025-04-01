package solution

import "sort"

func arrayPairSum1(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func arrayPairSum2(nums []int) int {
	counts := make([]int, 20001)
	for _, num := range nums {
		counts[num+10000]++
	}

	sum := 0
	isMinNum := true
	for num, count := range counts {
		for count > 0 {
			if isMinNum {
				sum += num - 10000
			}
			isMinNum = !isMinNum
			count--
		}
	}

	return sum
}
