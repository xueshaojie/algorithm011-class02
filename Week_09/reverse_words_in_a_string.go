// 两次翻转 时间复杂度O(n) 空间复杂度O(n)
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	l, r := 0, len(s)-1
	for l <= r && s[l] == ' ' {
		l++
	}
	for l <= r && s[r] == ' ' {
		r--
	}
	temp := []byte{}
	for l <= r {
		if s[l] == ' ' {
			start := l
			for start <= r && s[start] == ' ' {
				start++
			}
			temp = append(temp, ' ')
			l = start
		} else {
			temp = append(temp, s[l])
			l++
		}
	}

	for i, j := 0, len(temp)-1; i < len(temp)/2; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	for i := 0; i < len(temp); {
		start, end := i, i
		for end < len(temp) && temp[end] != ' ' {
			end++
		}
		end--
		i = end + 2
		for start < end {
			temp[start], temp[end] = temp[end], temp[start]
			start++
			end--
		}
	}
	return string(temp)
}