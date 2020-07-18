// dfs 时间复杂度O(MN) 空间复杂度O(MN)
func numIslands(grid [][]byte) int {
	var count int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// bfs 时间复杂度O(MN) 空间复杂度O(min(M,N))
ffunc numIslands(grid [][]byte) int {
    var count int 
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                grid[i][j] = '0'
                queue := [][2]int{{i,j}}
                for len(queue) > 0 {
                    first := queue[0]
                    queue = queue[1:]
                    x, y := first[0], first[1]
                    arr := [4][2]int{{x,y+1},{x-1,y},{x,y-1},{x+1,y}}
                    for k := 0; k < 4; k++ {
                        r, c := arr[k][0], arr[k][1]
                        if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == '1' {
                            grid[r][c] = '0'
                            queue = append(queue, [2]int{r, c})
                        }
                    }
                }
            }
        }
    }
    return count 
}