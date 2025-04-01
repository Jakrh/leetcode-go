package solution

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0, len(intervals)+1)

	idx := 0
	for idx < len(intervals) && intervals[idx][1] < newInterval[0] {
		result = append(result, intervals[idx])
		idx++
	}

	// Get start and end time of the new merged interval
	start, end := newInterval[0], newInterval[1]
	for idx < len(intervals) && intervals[idx][0] <= newInterval[1] {
		start = minInt(intervals[idx][0], start)
		end = maxInt(intervals[idx][1], end)
		idx++
	}
	result = append(result, []int{start, end})

	for idx < len(intervals) {
		result = append(result, intervals[idx])
		idx++
	}

	return result
}
