// 从中心往两侧延伸 时间复杂度O(n*n) 空间复杂度O(1)
func countSubstrings(s string) int {
	n, res := len(s), 0
	for i := 0; i <= 2*n-1; i++ {
		l := i / 2
		r := l + i%2
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}