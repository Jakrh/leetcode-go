package solution

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func numIslands1(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

// -----------

func destroyIsland(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	if i+1 < len(grid) {
		destroyIsland(grid, i+1, j)
	}
	if i-1 >= 0 {
		destroyIsland(grid, i-1, j)
	}
	if j+1 < len(grid[i]) {
		destroyIsland(grid, i, j+1)
	}
	if j-1 >= 0 {
		destroyIsland(grid, i, j-1)
	}
}

func numIslands2(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				destroyIsland(grid, i, j)
				count++
			}
		}
	}

	return count
}
