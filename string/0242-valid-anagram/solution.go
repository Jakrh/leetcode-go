package solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int, len([]rune(s)))
	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		if m[r] == 0 {
			return false
		}
		m[r]--
	}

	return true
}
