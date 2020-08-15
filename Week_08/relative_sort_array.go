import "sort"

// 计数排序 时间复杂度O(nlogn) 空间复杂度O(n)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, x := range arr1 {
		m[x]++
	}
	for _, x := range arr2 {
		for i := 0; i < m[x]; i++ {
			res = append(res, x)
		}
		delete(m, x)
	}
	temp := []int{}
	for k, v := range m {
		for v > 0 {
			temp = append(temp, k)
			v--
		}
	}
	sort.Ints(temp)
	res = append(res, temp...)
	return res
}