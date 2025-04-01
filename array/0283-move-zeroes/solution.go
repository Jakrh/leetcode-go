package solution

func moveZeroes1(nums []int) {
	// As let n = nums.length in JS
	n := len(nums)

	// As let zeroCount = 0 in JS
	zeroCount := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = nums[i]
		}
	}

	for i := n - 1; i >= n-zeroCount; i-- {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else if zeroCount > 0 {
			nums[i-zeroCount] = nums[i]
			nums[i] = 0
		}
	}
}
