package main

import (
	"os"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	// path := os.Args[1]
	// distance, _ := strconv.Atoi(os.Args[3])
	// word := os.Args[4]
	// trie := NewRadix()
	// str := readFile(path)
	// trie = addWordToTrie(&str, trie)

	// out := getwords(trie, word, distance)

	// serialize(trie, os.Args[2])
	deserialize(os.Args[2])
}