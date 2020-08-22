// 时间复杂度O(n) 空间复杂度O(n)
func toLowerCase(str string) string {
	strBytes := []byte(str)
	for i, s := range str {
		if strBytes[i] >= 'A' && strBytes[i] <= 'Z' {
			strBytes[i] = byte(s + 32)
		}
	}
	return string(strBytes)
}