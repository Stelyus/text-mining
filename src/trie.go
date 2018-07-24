package main

/**
 * For you my love FRANCK how to use the Trie
 * how to called :
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	dict       map[string]bool
	dictPrefix map[string]bool
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	d := make(map[string]bool, 1024)
	p := make(map[string]bool, 4096)
	p[""] = true
	return Trie{dict: d, dictPrefix: p}
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	t.dict[word] = true
	for i := 1; i < len(word); i++ {
		t.dictPrefix[word[:i]] = true
	}
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	return t.dict[word]
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.dict[prefix] || t.dictPrefix[prefix]
}
