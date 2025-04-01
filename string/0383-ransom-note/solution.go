package solution

const letterSize = 26

// Fastest took 9ms on LeetCode
func canConstruct1(ransomNote string, magazine string) bool {
	available := make(map[rune]int32, letterSize)

	for _, r := range magazine {
		available[r]++
	}
	for _, r := range ransomNote {
		if count, ok := available[r]; !ok || count == 0 {
			return false
		}
		available[r]--
	}

	return true
}

// Fastest took 0ms on LeetCode
func canConstruct2(ransomNote string, magazine string) bool {
	available := [letterSize]int32{}

	for _, r := range magazine {
		available[r-'a']++
	}
	for _, r := range ransomNote {
		if available[r-'a'] == 0 {
			return false
		}
		available[r-'a']--
	}

	return true
}
