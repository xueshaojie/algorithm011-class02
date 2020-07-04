import (
	"container/heap"
	"sort"
)

// sort 时间复杂度O(nlogn) 空间复杂度O(logn)
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// 堆 时间复杂度O(n) 空间复杂度O(k)
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	var h intHeap
	heap.Init(&h)
	for i := 0; i < len(arr); i++ {
		if i < k {
			heap.Push(&h, arr[i])
		} else {
			if arr[i] < (h)[0] {
				heap.Pop(&h)
				heap.Push(&h, arr[i])
			}
		}
	}
	return h
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}