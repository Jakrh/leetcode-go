package solution

import "strings"

func reverseString(s string) string {
	r := []rune(s)
	var sb strings.Builder
	for i := len(r) - 1; i >= 0; i-- {
		sb.WriteRune(r[i])
	}
	return sb.String()
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	sum := 0

	var sb strings.Builder
	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		i--
		j--
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		if sum&1 == 1 {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	if carry == 1 {
		sb.WriteRune('1')
	}

	return reverseString(sb.String())
}
