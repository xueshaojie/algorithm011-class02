type Trie struct {
	chrildren map[byte]*Trie
	isEnd     bool
}

func findWords(board [][]byte, words []string) []string {
	resMap := make(map[string]struct{}, 0)
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return res
	}
	trie := &Trie{chrildren: make(map[byte]*Trie, 0), isEnd: false}
	for _, word := range words {
		node := trie
		for i := 0; i < len(word); i++ {
			v, ok := node.chrildren[word[i]]
			if ok {
				node = v
			} else {
				new := &Trie{chrildren: make(map[byte]*Trie, 0), isEnd: false}
				node.chrildren[word[i]] = new
				node = new
			}
		}
		node.isEnd = true
	}
	m, n := len(board), len(board[0])
	dx, dy := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	var dfs func(i, j int, word string, node *Trie)
	dfs = func(i, j int, word string, node *Trie) {
		word += string(board[i][j])
		node = node.chrildren[board[i][j]]
		if node == nil {
			return
		}
		if node.isEnd {
			resMap[word] = struct{}{}
		}
		tmp := board[i][j]
		board[i][j] = '1'
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] != '1' {
				dfs(x, y, word, node)
			}
		}
		board[i][j] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _, ok := trie.chrildren[board[i][j]]; ok {
				dfs(i, j, "", trie)
			}
		}
	}
	for k, _ := range resMap {
		res = append(res, k)
	}
	return res
}
