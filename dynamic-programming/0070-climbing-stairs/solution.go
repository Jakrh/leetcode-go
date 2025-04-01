package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Memoization (top-down)
func climbStairs1(n int) int {
	return memorize(n, make([]int, n))
}

func memorize(n int, mem []int) int {
	if n == 1 || n == 2 {
		return n
	}

	if mem[n-1] == 0 {
		mem[n-1] = memorize(n-1, mem) + memorize(n-2, mem)
	}

	return mem[n-1]
}

// ------------------------------

// Tabulation (Bottom-up)
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		prev2, prev1 = prev1, prev2+prev1
	}

	return prev1
}
