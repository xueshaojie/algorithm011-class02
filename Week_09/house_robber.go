// dp 时间复杂度O(n) 空间复杂度O(n)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][1] = dp[i-1][0] + nums[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp 时间复杂度O(n) 空间复杂度O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}