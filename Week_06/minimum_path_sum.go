// dp 时间复杂度O(n*m) 空间复杂度O(m)
// 二维数组 dp[i][j] += min(dp[i+1][j], dp[i][j+1])
// 一维数组 dp[j] = grid[i][j] + min(dp[j], dp[j+1])
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := grid[m-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] += dp[i+1]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j == n-1 || dp[j] < dp[j+1] {
				dp[j] += grid[i][j]
			} else {
				dp[j] = grid[i][j] + dp[j+1]
			}
		}
	}
	return dp[0]
}