import "container/heap"

// 堆 时间复杂度O(nlogn) 空间复杂度O(n)
func nthUglyNumber(n int) int {
	m, res := make(map[int]struct{}), []int{}
	var h IntHeap // 然后定义最小堆
	heap.Init(&h)
	heap.Push(&h, 1)
	arr := []int{2, 3, 5}
	for i := 0; i < n; i++ {
		curr := heap.Pop(&h).(int)
		res = append(res, curr)
		for _, num := range arr {
			new := curr * num
			if _, ok := m[new]; !ok {
				m[new] = struct{}{}
				heap.Push(&h, new)
			}
		}
	}
	return res[n-1]
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// 动态规划 时间复杂度O(n) 空间复杂度O(n)
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

