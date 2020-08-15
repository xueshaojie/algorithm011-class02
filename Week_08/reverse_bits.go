// 逐位颠倒 时间复杂度O(1) 空间复杂度O(1)
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	power := 31
	for num > 0 {
		res += (num & 1) << power
		num = num >> 1
		power--
	}
	return res
}