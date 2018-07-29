package main

import "os"

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	trie := NewRadix()
	str := readFile(path)
	trie = addWordToTrie(&str, trie)

	
	// serialize(trie, os.Args[2])
}