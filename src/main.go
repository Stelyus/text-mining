package main

import (
		"os"
	"strconv"
	"fmt"
)


func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	distance, _ := strconv.Atoi(os.Args[3])
	word := os.Args[4]
	trie := NewRadix()
	str := readFile(path)
	trie = addWordToTrie(&str, trie)

	out := getwords(trie, word, distance)
	fmt.Println(out)
	//serialize(trie, os.Args[2])
}