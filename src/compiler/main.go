package main

import (
	"os"
	"radix"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	trie := radix.NewRadix()
	str := readFile(path)
	trie = addWordToTrie(&str, trie)

	serialize(trie, os.Args[2])
}