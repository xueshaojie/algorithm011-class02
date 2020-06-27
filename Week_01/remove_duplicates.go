// 双指针 时间复杂度O(n) 空间复杂度O(1)
func removeDuplicates(nums []int) int {
	j := 0
	for _, n := range nums {
		if n != nums[j] {
			j++
			nums[j] = n
		}
	}
	return j + 1
}
