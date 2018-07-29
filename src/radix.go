package main

import (
	"sort"
	"strings"
)

// WalkFn is used when walking the tree. Takes a
// Key and Value, returning if iteration should
// be terminated.
type WalkFn func(s string, v rune) bool

// leafNode is used to represent a Value
type leafNode struct {
	Key string
	Val rune
}

type node struct {
	// Leaf is used to store possible Leaf
	Leaf *leafNode

	// Prefix is the common Prefix we ignore
	Prefix string

	// Edges should be stored in-order for iteration.
	// We avoid a fully materialized slice to save memory,
	// since in most cases we expect to be sparse
	Edges edges
}

func (n *node) isLeaf() bool {
	return n.Leaf != nil
}

func (n *node) addEdge(newNode *node) {
	n.Edges = append(n.Edges, newNode)
	n.Edges.Sort()
}


func (n *node) replaceEdge(newNode *node) {
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

func (n *node) getEdge(label byte) *node {
	num := len(n.Edges)
	idx := sort.Search(num, func(i int) bool {
		return n.Edges[i].Prefix[0] >= label
	})
	if idx < num && n.Edges[idx].Prefix[0] == label {
		return n.Edges[idx]
	}

	return nil
}

type edges []*node

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i].Prefix[0] < e[j].Prefix[0]
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edges) Sort() {
	sort.Sort(e)
}

// Tree implements a radix tree. This can be treated as a
// Dictionary abstract data type. The main advantage over
// a standard hash map is Prefix-based lookups and
// ordered iteration,
type Tree struct {
	Root *node
}

// New returns an empty Tree
func NewRadix() *Tree {
	return NewFromMap(nil)
}

// NewFromMap returns a new tree containing the Keys
// from an existing map
func NewFromMap(m map[string]rune) *Tree {
	t := &Tree{Root: &node{}}
	for k, v := range m {
		t.Insert(k, v)
	}
	return t
}

// Len is used to return the number of elements in the tree

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

// Insert is used to add a newentry or update
// an existing entry. Returns if updated.
func (t *Tree) Insert(s string, v rune) (interface{}, bool) {
	var parent *node
	n := t.Root
	search := s
	for {
		// Handle Key exhaution
		if len(search) == 0 {
			if n.isLeaf() {
				old := n.Leaf.Val
				n.Leaf.Val = v
				return old, true
			}

			n.Leaf = &leafNode{
				Key: s,
				Val: v,
			}
			return nil, false
		}

		// Look for the edge
		parent = n
		n = n.getEdge(search[0])

		// No edge, create one
		if n == nil {

			// e := edge{
			// 	label: search[0],
			// 	node: &node{
			// 		Leaf: &leafNode{
			// 			Key: s,
			// 			Val: v,
			// 		},
			// 		Prefix: search,
			// 	},
			// }
			parent.addEdge(&node{
					Leaf: &leafNode{
						Key: s,
						Val: v,
					},
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

		// Split the node
		child := &node{
			Prefix: search[:commonPrefix],
		}

		// parent.replaceEdge(edge{
		// 	label: search[0],
		// 	node:  child,
		// })

		parent.replaceEdge(child)

		// Restore the existing node
		// child.addEdge(edge{
		// 	label: n.Prefix[commonPrefix],
		// 	node:  n,
		// })

		child.addEdge(n)
		n.Prefix = n.Prefix[commonPrefix:]

		// Create a new Leaf node
		Leaf := &leafNode{
			Key: s,
			Val: v,
		}

		// If the new Key is a subset, add to to this node
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.Leaf = Leaf
			return nil, false
		}

		// Create a new edge for the node
		// child.addEdge(edge{
		// 	label: search[0],
		// 	node: &node{
		// 		Leaf:   Leaf,
		// 		Prefix: search,
		// 	},
		// })

		child.addEdge(&node{
				Leaf:   Leaf,
				Prefix: search,
		})
		return nil, false
	}
}

// func (n *node) mergeChild() {
// 	e := n.Edges[0]
// 	child := e.node
// 	n.Prefix = n.Prefix + child.Prefix
// 	n.Leaf = child.Leaf
// 	n.Edges = child.Edges
// }

// Get is used to lookup a specific Key, returning
// the Value and if it was found
func (t *Tree) Get(s string) (interface{}, bool) {
	n := t.Root
	search := s
	for {
		// Check for Key exhaution
		if len(search) == 0 {
			if n.isLeaf() {
				return n.Leaf.Val, true
			}
			break
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search Prefix
		if strings.HasPrefix(search, n.Prefix) {
			search = search[len(n.Prefix):]
		} else {
			break
		}
	}
	return nil, false
}

// LongestPrefix is like Get, but instead of an
// exact match, it will return the longest Prefix match.
func (t *Tree) LongestPrefix(s string) (string, interface{}, bool) {
	var last *leafNode
	n := t.Root
	search := s
	for {
		// Look for a Leaf node
		if n.isLeaf() {
			last = n.Leaf
		}

		// Check for Key exhaution
		if len(search) == 0 {
			break
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search Prefix
		if strings.HasPrefix(search, n.Prefix) {
			search = search[len(n.Prefix):]
		} else {
			break
		}
	}
	if last != nil {
		return last.Key, last.Val, true
	}
	return "", nil, false
}

// Minimum is used to return the minimum Value in the tree
// func (t *Tree) Minimum() (string, interface{}, bool) {
// 	n := t.Root
// 	for {
// 		if n.isLeaf() {
// 			return n.Leaf.Key, n.Leaf.Val, true
// 		}
// 		if len(n.Edges) > 0 {
// 			n = n.Edges[0].node
// 		} else {
// 			break
// 		}
// 	}
// 	return "", nil, false
// }

// Maximum is used to return the maximum Value in the tree
// func (t *Tree) Maximum() (string, interface{}, bool) {
// 	n := t.Root
// 	for {
// 		if num := len(n.Edges); num > 0 {
// 			n = n.Edges[num-1].node
// 			continue
// 		}
// 		if n.isLeaf() {
// 			return n.Leaf.Key, n.Leaf.Val, true
// 		}
// 		break
// 	}
// 	return "", nil, false
// }

// Walk is used to walk the tree
func (t *Tree) Walk(fn WalkFn) {
	recursiveWalk(t.Root, fn)
}

// WalkPrefix is used to walk the tree under a Prefix
func (t *Tree) WalkPrefix(Prefix string, fn WalkFn) {
	n := t.Root
	search := Prefix
	for {
		// Check for Key exhaution
		if len(search) == 0 {
			recursiveWalk(n, fn)
			return
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search Prefix
		if strings.HasPrefix(search, n.Prefix) {
			search = search[len(n.Prefix):]

		} else if strings.HasPrefix(n.Prefix, search) {
			// Child may be under our search Prefix
			recursiveWalk(n, fn)
			return
		} else {
			break
		}
	}

}

// WalkPath is used to walk the tree, but only visiting nodes
// from the Root down to a given Leaf. Where WalkPrefix walks
// all the entries *under* the given Prefix, this walks the
// entries *above* the given Prefix.
func (t *Tree) WalkPath(path string, fn WalkFn) {
	n := t.Root
	search := path
	for {
		// Visit the Leaf Values if any
		if n.Leaf != nil && fn(n.Leaf.Key, n.Leaf.Val) {
			return
		}

		// Check for Key exhaution
		if len(search) == 0 {
			return
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			return
		}

		// Consume the search Prefix
		if strings.HasPrefix(search, n.Prefix) {
			search = search[len(n.Prefix):]
		} else {
			break
		}
	}
}

// recursiveWalk is used to do a pre-order walk of a node
// recursively. Returns true if the walk should be aborted
func recursiveWalk(n *node, fn WalkFn) bool {
	// Visit the Leaf Values if any
	if n.Leaf != nil && fn(n.Leaf.Key, n.Leaf.Val) {
		return true
	}

	// Recurse on the children
	for _, e := range n.Edges {
		if recursiveWalk(e, fn) {
			return true
		}
	}
	return false
}

// ToMap is used to walk the tree and convert it into a map
func (t *Tree) ToMap() map[string]rune {
	out := make(map[string]rune)
	t.Walk(func(k string, v rune) bool {
		out[k] = v
		return false
	})
	return out
}
