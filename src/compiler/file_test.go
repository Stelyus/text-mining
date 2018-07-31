package compiler

import (
	"testing"
	"radix"
)


func TestAddWordToTrie(test *testing.T) {
	path := "../../ressources/test.txt"
	str := ReadFile(path)

	trie := radix.NewRadix()
	AddWordToTrie(&str, trie)

	if len(trie.Root.Edges) != 7 {
		test.Errorf("Error len(trie.Root.Edges) got %d, expected %d", len(trie.Root.Edges), 7)
	}

	if trie.Root.Edges[0].Val != -1 {
		test.Errorf("Error trie.Root.Edges[0].Val got %d, expected %d", trie.Root.Edges[0].Val, -1)
	}

	if len(trie.Root.Edges[0].Edges) != 2 {
		test.Errorf("Error len(trie.Root.Edges[0].Edges) got %d, expected %d", len(trie.Root.Edges[0].Edges), 2)
	}

	if trie.Root.Edges[0].Prefix != "a" {
		test.Errorf("Error trie.Root.Edges[0].Prefix got %s, expected %s", trie.Root.Edges[0].Prefix, "a")

	}
}

func TestSerialize(test *testing.T) {
	path := "../../ressources/test.txt"
	str := ReadFile(path)

	trie := radix.NewRadix()
	AddWordToTrie(&str, trie)

	Serialize(trie, "../../test_ressources/dict_test.bin")
}