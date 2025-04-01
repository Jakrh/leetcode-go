package solution

// 001 001 001 001 001 001
// 001 002 003 004 005 006
// 001 003 006 010 015 021
// 001 004 010 020 035 056

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}
