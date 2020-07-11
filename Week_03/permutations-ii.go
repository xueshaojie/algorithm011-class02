import "sort"

// 回溯 时间复杂度O(n*n!) 空间复杂度O(n)
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	var dfs func(int)
	dfs = func(i int) {
		if i == n-1 {
			res = append(res, append([]int{}, nums...))
		}
		for k := i; k < n; k++ {
			if k != i && nums[k] == nums[i] {
				continue
			}
			nums[k], nums[i] = nums[i], nums[k]
			dfs(i + 1)
		}
		for k := n - 1; k > i; k-- {
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	dfs(0)

	return res
}