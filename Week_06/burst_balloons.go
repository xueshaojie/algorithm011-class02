// dp 时间复杂度O(n*n) 空间复杂度O(n*n)
func maxCoins(nums []int) int {
	n := len(nums)
	res := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		res[i] = make([]int, n+2)
	}
	vals := append([]int{1}, nums...)
	vals = append(vals, 1)

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				sum := vals[i] * vals[k] * vals[j]
				sum += res[i][k] + res[k][j]
				if sum > res[i][j] {
					res[i][j] = sum
				}
			}
		}
	}
	return res[0][n+1]
}