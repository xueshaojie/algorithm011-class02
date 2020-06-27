// 暴力求解 时间复杂度 O(k*n)
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		previous := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			previous, nums[j] = nums[j], previous
		}
	}
}

// 翻转 时间复杂度O(n) 空间复杂度O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}