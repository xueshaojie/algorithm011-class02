// 时间复杂度O(n) 空间复杂度O(1)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}