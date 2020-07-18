// 广度优先搜索 时间复杂度O(M*N) 空间复杂度O(M*N)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0
	beginQueue := []string{beginWord}
	beginUsed := make([]bool, len(wordList))
	wordMap := make(map[string]int)
	for i, word := range wordList {
		wordMap[word] = i
	}
	if _, ok := wordMap[endWord]; !ok {
		return step
	}
	for len(beginQueue) > 0 {
		step++
		l := len(beginQueue)
		for i := 0; i < l; i++ {
			chars := []byte(beginQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]
				for k := 'a'; k <= 'z'; k++ {
					chars[j] = byte(k)
					idx, ok := wordMap[string(chars)]
					if ok && !beginUsed[idx] {
						if string(chars) == endWord {
							return step + 1
						}
						beginQueue = append(beginQueue, string(chars))
						beginUsed[idx] = true
					}
				}
				chars[j] = o
			}
		}
		beginQueue = beginQueue[l:]
	}
	return 0
}

// 双向广度优先搜索 时间复杂度O(M*N) 空间复杂度O(M*N)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0
	beginQueue := []string{beginWord}
	beginUsed := make([]bool, len(wordList))
	endQueue := []string{endWord}
	endUsed := make([]bool, len(wordList))
	wordMap := make(map[string]int)

	for i, word := range wordList {
		wordMap[word] = i
	}

	if i, ok := wordMap[endWord]; ok {
		endUsed[i] = true
	} else {
		return step
	}
	for len(beginQueue) > 0 {
		step++
		l := len(beginQueue)

		for i := 0; i < l; i++ {
			chars := []byte(beginQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]
				for k := 'a'; k <= 'z'; k++ {
					chars[j] = byte(k)
					idx, ok := wordMap[string(chars)]
					if ok && !beginUsed[idx] {
						if endUsed[idx] {
							return step + 1
						}
						beginQueue = append(beginQueue, string(chars))
						beginUsed[idx] = true
					}
				}
				chars[j] = o
			}
		}
		beginQueue = beginQueue[l:]
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
			beginUsed, endUsed = endUsed, beginUsed
		}
	}

	return 0
}