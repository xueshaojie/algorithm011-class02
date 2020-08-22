// 直接反转 时间复杂度O(n) 空间复杂度O(n)
func reverseWords(s string) string {
	strBytes := []byte(s)
	i := 0
	for i < len(s) {
		start, end := i, i
		for end < len(s) && strBytes[end] != ' ' {
			end++
		}
		end--
		i = end + 2
		for start < end {
			strBytes[start], strBytes[end] = strBytes[end], strBytes[start]
			start++
			end--
		}
	}
	return string(strBytes)
}