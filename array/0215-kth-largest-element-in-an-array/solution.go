package solution

// count sort
func findKthLargest(nums []int, k int) int {
	// Find the max and min value in nums
	// The range of nums is [-1e4, 1e4]
	var maxVal, minVal int = 1e4, -1e4
	for _, num := range nums {
		maxVal = max(maxVal, num)
		minVal = min(minVal, num)
	}

	// Create a count array to store the frequency of each number
	counts := make([]int, maxVal-minVal+1)
	for _, num := range nums {
		counts[num-minVal]++
	}

	// Find the kth largest element
	for i := len(counts) - 1; i >= 0; i-- {
		k -= counts[i]
		if k <= 0 {
			return i + minVal
		}
	}

	return -1
}
