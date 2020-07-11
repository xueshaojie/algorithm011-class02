// 回溯 时间复杂度O(n*n!) 空间复杂度O(n)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(int)
	dfs = func(first int) {
		if first == len(nums) {
			res = append(res, append([]int{}, nums...))
		}
		for i := first; i < len(nums); i++ {
			nums[i], nums[first] = nums[first], nums[i]
			dfs(first + 1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	dfs(0)
	return res
} 
