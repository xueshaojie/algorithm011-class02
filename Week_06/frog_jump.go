// dp 时间复杂度O(n*n) 空间复杂度O(n*n)
func canCross(stones []int) bool {
	dp := make(map[int]map[int]int, len(stones))

	for _, n := range stones {
		dp[n] = make(map[int]int)
	}
	dp[0][0] = 0
	for i, n := range stones {
		for _, k := range dp[n] {
			for step := k - 1; step <= k+1; step++ {
				if _, ok := dp[n+step]; ok && step > 0 {
					dp[n+step][i] = step
				}
			}
		}
	}
	return len(dp[stones[len(stones)-1]]) != 0
}
