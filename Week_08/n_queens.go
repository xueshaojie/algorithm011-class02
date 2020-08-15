func solveNQueens(n int) [][]string {
	res := [][]int{}
	var dfs func(row, col, pie, na int, state []int)
	dfs = func(row, col, pie, na int, state []int) {
		if row >= n {
			res = append(res, append([]int{}, state...))
			return
		}
		bits := (^(col | pie | na)) & ((1 << n) - 1)
		for bits != 0 {
			p := bits & (-bits)
			i, temp := 0, (1 << uint(n-1))
			for p&temp == 0 {
				i++
				temp >>= 1
			}
			dfs(row+1, col|p, (pie|p)<<1, (na|p)>>1, append(state, i))
			bits &= bits - 1
		}
	}
	dfs(0, 0, 0, 0, []int{})
	return convertResult(res)
}

func convertResult(items [][]int) [][]string {
	res := [][]string{}
	for _, item := range items {
		temp := []string{}
		for _, x := range item {
			tmp := []byte{}
			for i := 0; i < len(item); i++ {
				tmp = append(tmp, '.')
			}
			tmp[x] = 'Q'
			temp = append(temp, string(tmp))
		}
		res = append(res, temp)
	}
	return res
}