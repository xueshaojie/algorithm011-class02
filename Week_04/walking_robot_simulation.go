type point struct {
	x int
	y int
}

// 模拟 时间复杂度O(N+K) 空间复杂度O(K)
func robotSim(commands []int, obstacles [][]int) int {
	points := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	obsMap := make(map[point]bool)
	for _, v := range obstacles {
		obsMap[point{v[0], v[1]}] = true
	}
	x, y, di, res := 0, 0, 0, 0
	for _, cmd := range commands {
		if cmd == -1 {
			di = (di + 3) % 4
		} else if cmd == -2 {
			di = (di + 1) % 4
		} else {
			for j := 0; j < cmd; j++ {
				if _, ok := obsMap[point{x + points[di][0], y + points[di][1]}]; ok {
					break
				} else {
					x, y = x+points[di][0], y+points[di][1]
				}
			}
			if res < x*x+y*y {
				res = x*x + y*y
			}
		}
	}
	return res
}