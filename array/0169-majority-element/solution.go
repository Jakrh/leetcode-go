package solution

// Answer 1
func majorityElement1(nums []int) int {
	halfN := len(nums) / 2
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
		if m[num] > halfN {
			return num
		}
	}

	return 0
}

// Answer 2
func majorityElement2(nums []int) int {
	majority, count := 0, 0

	for _, num := range nums {
		if num == majority {
			count++
		} else if count == 0 {
			majority, count = num, 1
		} else {
			count--
		}
	}

	return majority
}
