package solution

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	used := map[int]bool{}

	res := []int{}
	for _, num := range nums1 {
		m[num] = true
	}

	for _, num := range nums2 {
		if m[num] && !used[num] {
			res = append(res, num)
			used[num] = true
		}
	}

	return res
}
