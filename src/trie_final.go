package main

type TrieNode struct {
	val rune
	freq int
	children map[rune]*TrieNode
}


// Creates a new TrieNode as  root.
func NewTrie() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}


func (root *TrieNode) Add(key string, freq int) *TrieNode {
	runes := []rune(key)
	node := root
	for i := range runes {
		r := runes[i]
		if n, ok := node.children[r]; ok {
			node = n
		} else {
			if i == len(runes) - 1 {
				node  = node.NewNode(r, freq)
			} else{
				node = node.NewNode(r, -1)
			}
		}

	}
	return node
}

func (n TrieNode) Children() map[rune]*TrieNode {
	return n.children
}

func (root *TrieNode) Find(key string) (*TrieNode, bool) {
	node := findTrieNode(root, []rune(key))
	if node == nil {
		return nil, false
	}

	return node, true
}

// Creates and returns a pointer to a new child for the node.
func (n *TrieNode) NewNode(val rune, freq int) *TrieNode {
	node := &TrieNode{
		val:      val,
		freq: freq,
		children: make(map[rune]*TrieNode),
	}
	n.children[val] = node
	return node
}

func findTrieNode(node *TrieNode, runes []rune) *TrieNode {
	if node == nil {
		return nil
	}

	if len(runes) == 0 {
		return node
	}

	n, ok := node.children[runes[0]]
	if !ok {
		return nil
	}

	var nrunes []rune
	if len(runes) > 1 {
		nrunes = runes[1:]
	} else {
		nrunes = runes[0:0]
	}

	return findTrieNode(n, nrunes)
}