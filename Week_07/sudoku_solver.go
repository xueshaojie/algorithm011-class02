func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	var dfs func(board [][]byte) bool
	dfs = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for c := '1'; c <= '9'; c++ {
						if isValid(board, i, j, byte(c)) {
							board[i][j] = byte(c)
							if dfs(board) {
								return true
							} else {
								board[i][j] = '.'
							}
						}
					}
					return false
				}
			}
		}
		return true
	}
	dfs(board)
}

func isValid(board [][]byte, row, col int, char byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' && board[i][col] == char {
			return false
		}
		if board[row][i] != '.' && board[row][i] == char {
			return false
		}
		r := (row/3)*3 + i/3
		c := (col/3)*3 + i%3
		if board[r][c] != '.' && board[r][c] == char {
			return false
		}
	}
	return true
}