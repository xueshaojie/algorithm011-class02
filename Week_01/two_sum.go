// 暴力求解 时间复杂度O(n^2) 空间复杂度O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// map 时间复杂度O(n) 空间复杂度O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[n]; ok {
			return []int{v, i}
		}
		m[target-n] = i
	}
	return nil
}