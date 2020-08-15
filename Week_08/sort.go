// 选择排序 时间复杂度O(n^2) 空间复杂度O(1) 不稳定
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// 插入排序 时间复杂度O(n^2) 空间复杂度O(1) 稳定
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}

// 冒泡排序 时间复杂度O(n^2) 空间复杂度O(1) 稳定
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 快速排序 时间复杂度O(nlogn) 空间复杂度O(nlogn) 不稳定
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func partition(arr []int, left, right int) int {
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}

// 归并排序 时间复杂度O(nlogn) 空间复杂度O(n) 稳定
func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	for p := 0; p < len(temp); p++ {
		arr[left+p] = temp[p]
	}
}