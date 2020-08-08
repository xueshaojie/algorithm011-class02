// dfs 时间复杂度O(m*n) 空间复杂度O(m*n)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	dx, dy := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
				dfs(x, y)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// 并查集 时间复杂度O(m*n) 空间复杂度O(m*n)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	dx, dy := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	parent := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := n*i + j
			if grid[i][j] == '1' {
				parent[idx] = idx
			} else {
				parent[idx] = -1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				for k := 0; k < 4; k++ {
					x, y := i+dx[k], j+dy[k]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
						union(parent, i*n+j, x*n+y)
					}
				}
			}
		}
	}
	for i := 0; i < len(parent); i++ {
		if parent[i] == i {
			res++
		}
	}
	return res
}

func union(p []int, i, j int) {
	p1 := parent(p, i)
	p2 := parent(p, j)
	p[p1] = p2
}

func parent(p []int, i int) int {
	root := i
	for p[root] != root {
		root = p[root]
	}
	if p[i] != root {
		p[i] = root
	}
	return root
}