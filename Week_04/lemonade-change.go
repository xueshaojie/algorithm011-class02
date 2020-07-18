// 模拟情景 时间复杂度O(n) 空间复杂度O(1)
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, x := range bills {
		if x == 5 {
			five++
		} else if x == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}