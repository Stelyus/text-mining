// Package radix made for TEXT MINING, construct a radix tree
// that contains the trie structure and its relative function.
// Please see the tree structure, node structure and then edges structure
package radix

import (
	"sort"
)

// Tree implements a radix tree. This can be treated as a
// Dictionary abstract data type. The main advantage over
// a standard hash map is Prefix-based lookups and
// ordered iteration,
type Tree struct {
	Root *Node
}

// Node is the smallest struct in a Trie.
// It contains either a complete word or a part of the word and other relative information
type Node struct {
	// Prefix is the common Prefix from edges (eg. from its child)
	Prefix string

	// Val is the frequency of Key.
	// -1 if it's not a complete word
	Val rune

	// Key is the complete word
	// The behavior is undefined if it is not a complete word
	Key string

	// Edges should be stored in-order for iteration.
	// We avoid a fully materialized slice to save memory,
	// since in most cases we expect to be sparse.
	// It's a list containing the node's child. (type Edges = []*Node)
	Edges Edges
}



// IsLeaf check rather it is a leaf (a complete word)
func (n *Node) IsLeaf() bool {
	return n.Val != -1
}


// AddEdge add a edge (= child) into a node and then sort it
func (n *Node) AddEdge(newNode *Node) {
	n.Edges = append(n.Edges, newNode)
	n.Edges.Sort()
}


// replaceEdge replace a node by another node by checking its prefix
func (n *Node) replaceEdge(newNode *Node) {
	num := len(n.Edges)
	idx := sort.Search(num, func(i int) bool {
		return n.Edges[i].Prefix[0] >= newNode.Prefix[0]
	})
	if idx < num && n.Edges[idx].Prefix[0] == newNode.Prefix[0] {
		n.Edges[idx] = newNode
		return
	}

	panic("replacing missing edge")
}

// getEdges returns a pointer to a node containing the label
func (n *Node) getEdge(label byte) *Node {
	num := len(n.Edges)
	idx := sort.Search(num, func(i int) bool {
		return n.Edges[i].Prefix[0] >= label
	})
	if idx < num && n.Edges[idx].Prefix[0] == label {
		return n.Edges[idx]
	}

	return nil
}


// Edges is the list of childs of a node
type Edges []*Node

// Len returns the length of node's child
func (e Edges) Len() int {
	return len(e)
}

// Less return true if the first letter of i is smaller than the first letter of j
func (e Edges) Less(i, j int) bool {
	return e[i].Prefix[0] < e[j].Prefix[0]
}

// Swap edges at index i and j
func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Sort edges
func (e Edges) Sort() {
	sort.Sort(e)
}



// NewRadix returns an empty Tree
func NewRadix() *Tree {
	return &Tree{Root: &Node{
		Val: -1,
	}}
}

// longestPrefix finds the length of the shared Prefix
// of two strings
func longestPrefix(k1, k2 string) int {
	max := len(k1)
	if l := len(k2); l < max {
		max = l
	}
	var i int
	for i = 0; i < max; i++ {
		if k1[i] != k2[i] {
			break
		}
	}
	return i
}

// Insert is used to add a new entry or update
// an existing entry. Returns if updated.
func (t *Tree) Insert(s string, v rune) (interface{}, bool) {
	var parent *Node
	n := t.Root
	search := s
	for {
		// Handle Key exhaution
		if len(search) == 0 {
			if n.IsLeaf() {
				old := n.Val
				n.Val = v
				return old, true
			}

			n.Key = s
			n.Val = v
			return nil, false
		}

		// Look for the edge
		parent = n
		n = n.getEdge(search[0])

		// No edge, create one
		if n == nil {

			parent.AddEdge(&Node{
					Key: s,
					Val: v,
					Prefix: search,
			})
			return nil, false
		}

		// Determine longest Prefix of the search Key on match
		commonPrefix := longestPrefix(search, n.Prefix)
		if commonPrefix == len(n.Prefix) {
			search = search[commonPrefix:]
			continue
		}

		// Split the Node
		child := &Node{
			Prefix: search[:commonPrefix],
			Val: -1,
		}

		parent.replaceEdge(child)

		child.AddEdge(n)
		n.Prefix = n.Prefix[commonPrefix:]

		// If the new Key is a subset, add to to this Node
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.Key = s
			child.Val = v
			return nil, false
		}


		child.AddEdge(&Node{
				Key: s,
				Val: v,
				Prefix: search,
		})
		return nil, false
	}
}