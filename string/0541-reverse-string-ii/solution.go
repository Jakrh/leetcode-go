package solution

func reverseStr(s string, k int) string {
	b := []byte(s)
	n := len(b)
	for base := 0; base < n; base += 2 * k {
		l, r := base, min(base+k-1, n-1)
		for l <= r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return string(b)
}
