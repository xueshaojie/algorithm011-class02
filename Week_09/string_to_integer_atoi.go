import (
	"math"
	"strings"
)

// 时间复杂度O(n) 空间复杂度O(1)
func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	sign = 1
	if s[0] >= '0' && s[0] <= '9' {
		abs = s
	} else if s[0] == '+' {
		abs = s[1:]
	} else if s[0] == '-' {
		sign = -1
		abs = s[1:]
	} else {
		return
	}
	for i, b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}
	return
}

func convert(sign int, absStr string) int {
	res := 0
	for _, b := range absStr {
		res = res*10 + int(b-'0')
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign == -1 && res*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * res
}