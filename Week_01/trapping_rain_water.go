// 暴力法 时间复杂度O(n^2) 空间复杂度O(1)
func trap(height []int) int {
	res := 0
	for i, n := range height {
		leftMax, rightMax, min := 0, 0, 0
		for j := 0; j <= i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}

		for j := i; j < len(height); j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}
		if leftMax > rightMax {
			min = rightMax
		} else {
			min = leftMax
		}
		res += min - n
	}
	return res
}

// 动态编程 时间复杂度O(n) 空间复杂度O(n)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax, rightMax, res := make([]int, len(height)), make([]int, len(height)), 0
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i, _ := range height {
		var min int
		if leftMax[i] > rightMax[i] {
			min = rightMax[i]
		} else {
			min = leftMax[i]
		}
		res += min - height[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 双指针法 时间复杂度O(n) 空间复杂度O(1)
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, res := 0, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}