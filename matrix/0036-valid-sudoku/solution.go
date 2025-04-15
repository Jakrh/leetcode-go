package solution

func isValidSudoku(board [][]byte) bool {
	size := 9
	usedCol := make([][]bool, size)
	usedRow := make([][]bool, size)
	usedBlock := make([][]bool, size)
	for i := 0; i < size; i++ {
		usedCol[i] = make([]bool, size)
		usedRow[i] = make([]bool, size)
		usedBlock[i] = make([]bool, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0' - 1
				block := i/3*3 + j/3
				if usedCol[i][num] || usedRow[j][num] || usedBlock[block][num] {
					return false
				}
				usedCol[i][num], usedRow[j][num], usedBlock[block][num] = true, true, true
			}
		}
	}

	return true
}
