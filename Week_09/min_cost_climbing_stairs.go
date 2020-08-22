// 时间复杂度O(n) 空间复杂度O(1)
func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f1, f2 = cost[i]+min(f1, f2), f1
	}
	return min(f1, f2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}