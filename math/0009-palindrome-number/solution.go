package solution

import "strconv"

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	for i := 0; i < len(xStr)>>1; i++ {
		if xStr[i] != xStr[len(xStr)-i-1] {
			return false
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	tmp := x
	reversedX := 0
	for tmp > 0 {
		reversedX = reversedX*10 + tmp%10
		tmp /= 10
	}

	return x == reversedX
}
