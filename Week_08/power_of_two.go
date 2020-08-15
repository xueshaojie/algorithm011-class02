// 位操作 时间复杂度O(1) 空间复杂度O(1)
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
