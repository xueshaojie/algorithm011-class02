// 时间复杂度O(n) 空间复杂度O(1)
func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
