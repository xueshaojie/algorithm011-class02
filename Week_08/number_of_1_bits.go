// 位操作 取掉最低位1 时间复杂度O(1) 空间复杂度O(1)
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		num &= num - 1
		res++
	}
	return res
}