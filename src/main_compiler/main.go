package main

import (
	"os"
	"radix"
	"compiler"
)

func main() {
	path := os.Args[1]
	trie := radix.NewRadix()
	str := compiler.ReadFile(path)
	trie = compiler.AddWordToTrie(&str, trie)

	compiler.Serialize(trie, os.Args[2])
}