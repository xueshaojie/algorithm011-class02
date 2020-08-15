// 归并排序 时间复杂的O(nlogn) 空间复杂度O(n)
func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	temp := make([]int, r-l+1)
	mid := l + (r-l)/2
	count := mergeSort(nums, l, mid) + mergeSort(nums, mid+1, r)
	i, t, k := l, l, 0
	for j := mid + 1; j <= r; j++ {
		for i <= mid && nums[i] <= 2*nums[j] {
			i++
		}
		for t <= mid && nums[t] < nums[j] {
			temp[k] = nums[t]
			t++
			k++
		}
		temp[k] = nums[j]
		k++
		count += mid - i + 1

	}
	for t <= mid {
		temp[k] = nums[t]
		t++
		k++
	}
	for p := 0; p < len(temp); p++ {
		nums[l+p] = temp[p]
	}
	return count
} 