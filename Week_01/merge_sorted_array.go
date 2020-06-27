import "sort"

// 时间复杂度O((n+m)log(n + m)) 空间复杂度O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], num2[:n]...)
	sort.Ints(nums1)
}

// 双指针从后往前 时间复杂度O(n+m) 空间复杂度O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}