package solution

func longestPalindrome(s string) int {
	m := map[byte]bool{}
	counter := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, ok := m[b]; ok {
			delete(m, b)
			counter++
		} else {
			m[b] = true
		}
	}
	if len(m) == 0 {
		return 2 * counter
	} else {
		return 2*counter + 1
	}
}
