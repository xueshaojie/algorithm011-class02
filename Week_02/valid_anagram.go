import "sort"

// map 时间复杂度O(n) 空间复杂度O(1)
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if v, _ := m[r]; v == 1 {
			delete(m, r)
			continue
		}
		m[r]--
	}
	return len(m) == 0
}

// sort 时间复杂度O(nlogn) 空间复杂度O(n)
func isAnagram(s string, t string) bool {
	sRune, tRune := mySort(s), mySort(t)
	if len(sRune) != len(tRune) {
		return false
	}
	sort.Sort(sRune)
	sort.Sort(tRune)
	for i := 0; i < len(sRune); i++ {
		if sRune[i] != tRune[i] {
			return false
		}
	}
	return true
}

type mySort []rune

func (m mySort) Len() int {
	return len(m)
}

func (m mySort) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m mySort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}