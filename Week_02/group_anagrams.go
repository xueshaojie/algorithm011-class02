import "sort"

// map key为排序后的字符串 时间复杂度O(NKlogK) 空间复杂度(KlogK)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := [][]string{}
	for _, str := range strs {
		s := runes(str)
		sort.Sort(s)
		k := string(s)
		if i, ok := m[k]; ok {
			res[i] = append(res[i], str)
		} else {
			m[k] = len(res)
			res = append(res, []string{str})
		}
	}
	return res
}

type runes []rune

func (r runes) Len() int {
	return len(r)
}
func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}
func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}