import "sort"

// 时间复杂度 O(nlogn) 空间复杂度O(logn) n为区间数量
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else if interval[1] > res[len(res)-1][1] {
			res[len(res)-1][1] = interval[1]
		}
	}
	return res
}