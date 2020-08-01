import "sort"

// 时间复杂度O(M) 空间复杂度O(1)
func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		arr[tasks[i]-'A']++
	}
	sort.Ints(arr)
	maxVal := arr[25] - 1
	idleSlots := maxVal * n
	for i := 24; i >= 0 && arr[i] > 0; i-- {
		if arr[i] > maxVal {
			idleSlots -= maxVal
		} else {
			idleSlots -= arr[i]
		}
	}
	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}
