package solution

// shifts,      n, n as binary
//  0,          1,                               1
//  2,          4,                             100
//  4,         16,                           10000
//  6,         64,                         1000000
//  8,        256,                       100000000
// 10,       1024,                     10000000000
// 12,       4096,                   1000000000000
// 14,      16384,                 100000000000000
// 16,      65536,               10000000000000000
// 18,     262144,             1000000000000000000
// 20,    1048576,           100000000000000000000
// 22,    4194304,         10000000000000000000000
// 24,   16777216,       1000000000000000000000000
// 26,   67108864,     100000000000000000000000000
// 28,  268435456,   10000000000000000000000000000
// 30, 1073741824, 1000000000000000000000000000000

// time: 0ms
func isPowerOfFour1(n int) bool {
	if n < 0 {
		return false
	}

	const maxShift = 30
	shift := 0
	for {
		if n == 1<<shift {
			return true
		}

		if shift > maxShift {
			return false
		}
		shift += 2
	}
}

// time: 0ms
func isPowerOfFour2(n int) bool {
	return n&0x55555555 > 0 && n&(n-1) == 0

	// 1. Ensure at least 1 bit one in correct position
	//    0x5555555 -> 0b1010101010101010101010101010101
	// 2. Ensure there is the only 1 bit one
}
