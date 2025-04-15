package solution

const size = 9

func isValid(board [][]byte, row, col int, target byte) bool {
	blockRow, blockCol := 3*(row/3), 3*(col/3)
	for i := 0; i < size; i++ {
		if board[i][col] == target {
			return false
		}
		if board[row][i] == target {
			return false
		}
		if board[blockRow+i/3][blockCol+i%3] == target {
			return false
		}
	}

	return true
}

func solve(board [][]byte, row, col int) bool {
	for i := row; i < size; i++ {
		for j := col; j < size; j++ {
			if board[i][j] != '.' {
				continue
			}
			for num := '1'; num <= '9'; num++ {
				if isValid(board, i, j, byte(num)) {
					board[i][j] = byte(num)
					if solve(board, i, j+1) {
						return true
					}
					board[i][j] = '.'
				}
			}

			return false
		}
		col = 0
	}

	return true
}

func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}
