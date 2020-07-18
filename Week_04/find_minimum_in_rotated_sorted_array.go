// 二分查找 时间复杂度O(logn) 空间复杂度O(1)
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l < r {
		mid = (r + l) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}