package solution

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if value, ok := numMap[complement]; ok {
			return []int{value, i}
		} else {
			numMap[nums[i]] = i
		}
	}

	return []int{}
}
