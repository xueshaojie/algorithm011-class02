// 时间复杂度O(n) 空间复杂度O(n)
func numJewelsInStones(J string, S string) int {
	res := 0
	m := make(map[rune]int)
	for _, j := range J {
		m[j] = 1
	}
	for _, s := range S {
		if m[s] == 1 {
			res++
		}
	}
	return res
}