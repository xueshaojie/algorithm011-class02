// 二分查找 时间复杂度O(logN) 空间复杂度O(1)
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}