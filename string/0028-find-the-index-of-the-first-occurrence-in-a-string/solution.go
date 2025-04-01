package solution

// 0 ms
func strStr1(haystack string, needle string) int {
	res := 0

	i := 0
	for i < len(haystack) {
		if haystack[i] == needle[0] {
			res = i

			i++
			j := 1
			retryStartIdx := -1
			for j < len(needle) && i < len(haystack) {
				if retryStartIdx == -1 && haystack[i] == needle[0] {
					retryStartIdx = i
				}
				if haystack[i] != needle[j] {
					if retryStartIdx != -1 {
						i = retryStartIdx
					}
					break
				}
				i++
				j++
			}
			if j == len(needle) {
				return res
			}
		} else {
			i++
		}
	}

	return -1
}

// 0 ms
func strStr2(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i <= len(haystack)-n; i++ {
		if haystack[i] == needle[0] && haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
