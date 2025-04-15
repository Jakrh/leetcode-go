package solution

func checkIfExist(arr []int) bool {
	m := map[int]bool{}

	for _, num := range arr {
		if m[num*2] || num&1 == 0 && m[num/2] {
			return true
		}
		m[num] = true
	}

	return false
}
