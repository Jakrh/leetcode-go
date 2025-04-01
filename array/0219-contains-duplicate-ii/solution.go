package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[num]; ok && i-j <= k {
			return true
		}
		m[num] = i
	}

	return false
}
