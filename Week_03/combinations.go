// 回溯 时间复杂度 空间复杂度
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	curr := make([]int, 0)

	var backtrack func(int)
	backtrack = func(first int) {
		if len(curr) == k {
			res = append(res, append([]int{}, curr...))
		}
		for i := first; i <= n; i++ {
			curr = append(curr, i)
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1)

	return res
}

// 字典序 时间复杂度 空间复杂度
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var nums []int
	for i := 1; i <= k; i++ {
		nums = append(nums, i)
	}
	nums = append(nums, n+1)
	j := 0
	for j < k {
		res = append(res, append([]int{}, nums[:k]...))
		for j = 0; j < k && nums[j+1] == nums[j]+1; j++ {
			nums[j] = j + 1
		}
		nums[j] += 1
	}
	return res
}