package solution

func lengthOfLongestSubstring(s string) int {
	m := map[rune]bool{}

	count := 0
	longest := 0
	for _, r := range s {
		if !m[r] {
			count++
			m[r] = true
		} else {
			count = 1
		}
		if count > longest {
			longest = count
		}
	}

	return longest
}
