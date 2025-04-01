package solution

import "strings"

func repeatedSubstringPattern1(s string) bool {
	ss := (s + s)[1 : 2*len(s)-1]

	return strings.Contains(ss, s)
}

func repeatedSubstringPattern2(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			isFactor := true
			for j := 0; j < len(s); j += i {
				if s[:i] != s[j:j+i] {
					isFactor = false
					break
				}
			}
			if isFactor {
				return true
			}
		}
	}

	return false
}
