// 时间复杂度O(n) 空间复杂度O(1)
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, x := range s {
		m[x]++
	}
	for _, x := range t {
		if m[x] == 1 {
			delete(m, x)
		} else {
			m[x]--
		}
	}
	return len(m) == 0
}
