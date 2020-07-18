// dfs 时间复杂度O(M*N) 空间复杂度O(M*N)
func updateBoard(board [][]byte, click []int) [][]byte {
	var d = [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	a, b := click[0], click[1]
	if board[a][b] == 'M' {
		board[a][b] = 'X'
	} else if board[a][b] == 'E' {
		m, n := len(board), len(board[0])
		var dfs func(int, int)
		dfs = func(i, j int) {
			c := byte('0')
			for _, di := range d {
				x, y := i+di[0], j+di[1]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'M' {
					c++
				}
			}
			if c > '0' {
				board[i][j] = c
			} else {
				board[i][j] = 'B'
				for _, di := range d {
					x, y := i+di[0], j+di[1]
					if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'E' {
						dfs(x, y)
					}
				}
			}
		}
		dfs(a, b)
	}
	return board
}

// bfs 时间复杂度O(M*N) 空间复杂度O(M*N)
type idx struct{ x, y int }

var d = [8]idx{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	a, b := click[0], click[1]
	if board[a][b] == 'M' {
		board[a][b] = 'X'
	} else if board[a][b] == 'E' {
		q, m, n := map[idx]bool{{a, b}: false}, len(board), len(board[0])
		for len(q) > 0 {
			p := map[idx]bool{}
			for i := range q {
				c, t := byte('0'), []idx{}
				for _, di := range d {
					x, y := i.x+di.x, i.y+di.y
					if 0 <= x && x < m && 0 <= y && y < n {
						if board[x][y] == 'M' {
							c++
						}
						if board[x][y] == 'E' {
							t = append(t, idx{x, y})
						}
					}
				}
				if c > '0' {
					board[i.x][i.y] = c
				} else {
					board[i.x][i.y] = 'B'
					for _, j := range t {
						p[j] = false
					}
				}
			}
			q = p
		}
	}
	return board
}
