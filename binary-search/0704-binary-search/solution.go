package solution

func search(nums []int, target int) int {
	l, m, r := 0, 0, len(nums)-1
	for l <= r {
		m = (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
