package solution

func getSumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		unitsDigit := n % 10
		sum += unitsDigit * unitsDigit
		n /= 10
	}

	return sum
}

func isHappy1(n int) bool {
	slow, fast := n, n
	for {
		slow = getSumOfSquares(slow)
		if slow == 1 {
			return true
		}

		fast = getSumOfSquares(getSumOfSquares(fast))
		if slow == fast {
			return false
		}
	}
}

// ----------

func isHappy2(n int) bool {
	for n > 6 {
		v := n
		s := []int{}
		for v > 0 {
			s = append(s, v%10)
			v /= 10
		}
		for i, elem := range s {
			if i == 0 {
				n = 0
			}
			n += elem * elem
		}
	}

	return n == 1
}
