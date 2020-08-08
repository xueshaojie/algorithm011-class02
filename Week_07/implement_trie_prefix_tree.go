type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{children: make(map[byte]*Trie, 0), isEnd: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		v, ok := node.children[word[i]]
		if ok {
			node = v
		} else {
			trie := Constructor()
			node.children[word[i]] = &trie
			node = &trie
		}
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		v, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = v
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		v, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = v
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */