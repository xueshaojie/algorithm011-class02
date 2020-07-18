// 二分查找 时间复杂度O(logn) 空间复杂度O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	l, r, mid, x := 0, m*n-1, 0, 0
	for l <= r {
		mid = (l + r) / 2
		x = matrix[mid/n][mid%n]
		if x == target {
			return true
		} else if x > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}