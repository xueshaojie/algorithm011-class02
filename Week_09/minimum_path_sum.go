// dp 时间复杂度O(mn) 空间复杂度O(mn)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := grid
	m, n := len(dp), len(dp[0])
	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] += dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] += min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}