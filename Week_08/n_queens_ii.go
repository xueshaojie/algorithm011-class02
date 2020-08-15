
func totalNQueens(n int) int {
	count := 0
	var dfs func(row, col, pie, na int)
	dfs = func(row, col, pie, na int) {
		if row >= n {
			count++
			return
		}
		bits := (^(col | pie | na)) & ((1 << n) - 1)
		for bits > 0 {
			p := bits & -bits
			bits &= bits - 1
			dfs(row+1, col|p, (pie|p)<<1, (na|p)>>1)
		}
	}
	dfs(0, 0, 0, 0)
	return count
}