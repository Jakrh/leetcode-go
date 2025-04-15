package solution

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if ok := numMap[num]; ok {
			return true
		}
		numMap[num] = true
	}

	return false
}
