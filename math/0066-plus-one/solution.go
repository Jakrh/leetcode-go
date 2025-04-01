package solution

// Answer 1
func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

// Answer 2
func plusOne2(digits []int) []int {
	carry := 0
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
	} else {
		digits[len(digits)-1] = 0
		carry = 1
		for i := len(digits) - 2; i >= 0; i-- {
			sum := digits[i] + carry
			if sum < 10 {
				digits[i] = sum
				carry = 0
				break
			} else {
				digits[i] = sum - 10
			}
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
