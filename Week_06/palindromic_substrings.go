// dp 时间复杂度O(n*n) 空间复杂度O(n*n)
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}
