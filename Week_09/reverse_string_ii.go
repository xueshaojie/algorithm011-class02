// 时间复杂度O(n) 空间复杂度O(1)
func reverseStr(s string, k int) string {
	strBytes := []byte(s)
	for start := 0; start < len(s); start += 2 * k {
		end := min(start+k-1, len(s)-1)
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
		}
	}
	return string(strBytes)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}