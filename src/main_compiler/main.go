package main

import (
	"os"
	"radix"
	"compiler"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	trie := radix.NewRadix()
	str := compiler.ReadFile(path)
	trie = compiler.AddWordToTrie(&str, trie)

	compiler.Serialize(trie, os.Args[2])
}