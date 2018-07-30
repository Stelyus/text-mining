package radix

import (
	"sort"
	"strings"
)

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// WalkFn is used when walking the tree. Takes a
// Key and Value, returning if iteration should
// be terminated.
type WalkFn func(s string, v rune) bool

type Node struct {
	// Leaf is used to store possible Leaf
	// Leaf *leafNode

	// Prefix is the common Prefix we ignore
	Prefix string
	Val rune
	Key string

	// Edges should be stored in-order for iteration.
	// We avoid a fully materialized slice to save memory,
	// since in most cases we expect to be sparse
	Edges edges
}

func (n *Node) IsLeaf() bool {
	return n.Val != -1
}

func (n *Node) AddEdge(newNode *Node) {
	n.Edges = append(n.Edges, newNode)
	n.Edges.Sort()
}


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

type edges []*Node

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
	Root *Node
}

// New returns an empty Tree
func NewRadix() *Tree {
	return NewFromMap(nil)
}

// NewFromMap returns a new tree containing the Keys
// from an existing map
func NewFromMap(m map[string]rune) *Tree {
	t := &Tree{Root: &Node{
		Val: -1,
	}}
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

		// Create a new Leaf Node
		
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

// Get is used to lookup a specific Key, returning
// the Value and if it was found
func (t *Tree) Get(s string) (interface{}, bool) {
	n := t.Root
	search := s
	for {
		// Check for Key exhaution
		if len(search) == 0 {
			if n.IsLeaf() {
				return n.Val, true
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

// Walk is used to walk the tree
func (t *Tree) Walk(fn WalkFn) {
	recursiveWalk(t.Root, fn)
}

// recursiveWalk is used to do a pre-order walk of a Node
// recursively. Returns true if the walk should be aborted
func recursiveWalk(n *Node, fn WalkFn) bool {
	// Visit the Leaf Values if any
	if n.IsLeaf() && fn(n.Key, n.Val) {
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
