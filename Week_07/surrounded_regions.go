// dfs 时间复杂度O(m*n) 空间复杂度O(m*n)
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	dx, dy := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		board[i][j] = '#'
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
				dfs(x, y)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}