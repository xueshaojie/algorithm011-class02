import "container/heap"

// 堆 时间复杂度O(nlogk) 空间复杂度O(n)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	var pq PriorityQueue // 然后定义最小堆
	heap.Init(&pq)
	i := 0
	for key, v := range m {
		if i < k {
			heap.Push(&pq, &Item{
				value:    key,
				priority: v,
			})
			i++
		} else {
			if v > (pq)[0].priority {
				heap.Pop(&pq)
				heap.Push(&pq, &Item{
					value:    key,
					priority: v,
				})
			}
		}
	}
	var res []int
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.value)
	}
	return res
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}