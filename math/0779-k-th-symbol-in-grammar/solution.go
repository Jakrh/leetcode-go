package solution

import "math/bits"

func kthGrammar(n int, k int) int {
	return bits.OnesCount32(uint32(k)-1) & 1
}
