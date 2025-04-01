package solution

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	currentArea := 0
	result := 0

	for l < r {
		currentArea = (r - l) * minInt(height[l], height[r])
		if currentArea > result {
			result = currentArea
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return result
}
