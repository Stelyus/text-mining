package main

import (
	"sort"
	"strings"
	"fmt"
)

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

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

func (n *node) isLeaf() bool {
	return n.Val != -1
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
	t := &Tree{Root: &node{
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
	var parent *node
	n := t.Root
	search := s
	for {
		// Handle Key exhaution
		if len(search) == 0 {
			if n.isLeaf() {
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

			parent.addEdge(&node{
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

		// Split the node
		child := &node{
			Prefix: search[:commonPrefix],
			Val: -1,
		}

		parent.replaceEdge(child)

		child.addEdge(n)
		n.Prefix = n.Prefix[commonPrefix:]

		// Create a new Leaf node
		
		// If the new Key is a subset, add to to this node
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.Key = s
			child.Val = v
			return nil, false
		}


		child.addEdge(&node{
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
			if n.isLeaf() {
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

// recursiveWalk is used to do a pre-order walk of a node
// recursively. Returns true if the walk should be aborted
func recursiveWalk(n *node, fn WalkFn) bool {
	// Visit the Leaf Values if any
	if n.isLeaf() && fn(n.Key, n.Val) {
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


func checkDamerau(ref string, currentWord string, distance int, isleaf bool) bool{

	//fmt.Println(ref, currentWord, distance)
	// check if length of the current word doesn't go way over distance
	if len(currentWord) - len(ref) > distance {
		return true
	}

	if currentWord != "" {
		d := 0
		// get the length of the current prefix
		mini := min(len(ref), len(currentWord))

		if len(ref) > mini {
			// check the prefix of the ref and the suffix to correctly check for damerau leveisntein distance
			prefix := ref[:mini]
			suffix := ref[1 : mini+1]
			if strings.Contains(currentWord, "nai") {
				fmt.Println("prefixe2:", prefix, suffix, currentWord)

			}
			d = min(DamerauLevenshtein(prefix, currentWord), DamerauLevenshtein(suffix, currentWord))
		} else {
			currentWord = currentWord[:mini]
			d = DamerauLevenshtein(ref, currentWord)
		}

		d = min(d, DamerauLevenshtein(ref, currentWord))
		if strings.Contains(currentWord, "nai") {

		fmt.Println("prefixe:", ref, currentWord)
		fmt.Println("distance is:", d)
		fmt.Println("---")
		}
		return !(d <= distance)
	}
	return false
}


