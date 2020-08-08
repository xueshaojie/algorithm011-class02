// 迭代 时间复杂度O(1) 空间复杂度O(1)
func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[int]int, 9)
	cols := make(map[int]map[int]int, 9)
	boxs := make(map[int]map[int]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]int, 9)
		cols[i] = make(map[int]int, 9)
		boxs[i] = make(map[int]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := int(board[i][j])
			index := (i/3)*3 + j/3
			rows[i][n]++
			cols[j][n]++
			boxs[index][n]++
			if rows[i][n] > 1 || cols[j][n] > 1 || boxs[index][n] > 1 {
				return false
			}
		}
	}
	return true
}