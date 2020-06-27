
// 双指针法 时间复杂度O(n) 空间复杂度O(1)
func moveZeroes(nums []int) {
	j := 0
	for _, n := range nums {
		if n != 0 {
			nums[j] = n
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

// 双指针法直接交换两个值 时间复杂度O(n) 空间复杂度O(1)
func moveZeroes(nums []int) {
	j := 0
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[j] = 0, n
			j++
		}
	}
}