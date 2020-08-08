// dfs 时间复杂度O(n*n) 空间复杂度O(n)
func findCircleNum(M [][]int) int {
	res := 0
	n := len(M)
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			res++
			dfs(i)
		}
	}
	return res
}

// 并查集 时间复杂度O(n*n) 空间复杂度O(n)
func findCircleNum(M [][]int) int {
	res := 0
	n := len(M)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && i != j {
				union(parent, i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if parent[i] == i {
			res++
		}
	}
	return res
}

func union(p []int, i, j int) {
	p1 := parent(p, i)
	p2 := parent(p, j)
	p[p1] = p2
}

func parent(p []int, i int) int {
	root := i
	for p[root] != root {
		root = p[root]
	}
	if p[i] != i {
		p[i] = root
	}
	return root
}
