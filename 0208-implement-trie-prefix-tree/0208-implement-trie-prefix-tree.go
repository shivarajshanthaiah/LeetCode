type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &Trie{children: make(map[rune]*Trie)}
		}
		node = node.children[char]
	}
	node.isEnd = true

}

func (this *Trie) Search(word string) bool {
	node := this
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, char := range prefix {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
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