import "unicode"

// 时间复杂度O(n) 空间复杂度O(1)
func reverseOnlyLetters(S string) string {
	res := ""
	j := len(S) - 1
	for _, s := range S {
		if unicode.IsLetter(s) {
			for !unicode.IsLetter(rune(S[j])) {
				j--
			}
			res += string(S[j])
			j--
		} else {
			res += string(s)
		}
	}
	return res
}