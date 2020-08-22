// 时间复杂度O(n) 空间复杂度O(1)
func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, str := range s {
		m[str]++
	}
	for i, str := range s {
		if m[str] == 1 {
			return i
		}
	}
	return -1
}