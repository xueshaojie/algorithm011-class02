// 遍历 时间复杂度O(n) 空间复杂度O(1)
func jump(nums []int) int {
	end, maxPos, step := 0, 0, 0
	for i, n := range nums {
		if i == len(nums)-1 {
			continue
		}
		if maxPos < i+n {
			maxPos = i + n
		}
		if end == i {
			end = maxPos
			step++
		}
	}
	return step
}
