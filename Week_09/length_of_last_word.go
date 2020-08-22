// 时间复杂度O(n) 空间复杂度O(1)
func lengthOfLastWord(s string) int {
	res := 0
	end := len(s) - 1
	for end > 0 && s[end] == ' ' {
		end--
	}
	for j := end; j >= 0; j-- {
		if s[j] == ' ' {
			break
		}
		res++
	}
	return res
}