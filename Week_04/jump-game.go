// 贪心 时间复杂度O(N) 空间复杂度O(1)
func canJump(nums []int) bool {
	rightMost := 0
	for i, n := range nums {
		if rightMost >= i && rightMost < i+n {
			rightMost = i + n
		}
		if rightMost >= len(nums)-1 {
			return true
		}
	}
	return false
}

