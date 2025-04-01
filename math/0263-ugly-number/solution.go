package solution

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, x := range []int{3, 5} {
		for n%x == 0 {
			n /= x
		}
	}

	// Check if n is power of 2
	return n&(n-1) == 0
}
