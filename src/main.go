package main

import (
	"os"
	// "fmt"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	// distance, _ := strconv.Atoi(os.Args[3])
	// word := os.Args[4]
	trie := NewRadix()
	str := readFile(path)
	trie = addWordToTrie(&str, trie)

	// out := getwords(trie, word, distance)
	// out := getwords(trie, word, distance)
	//fmt.Println(DamerauLevenshtein(word, "abilo"), distance)
	//fmt.Println(out)

	// for k,v := range out{
	// 	fmt.Println("{\"word\": \"", k, "\"freq\":", v, "\"distance\":", 2)
	// }
	// fmt.Println(len(out))
	// fmt.Println(trie.Root.Edges[3].Edges[1].Prefix)
	// fmt.Println(trie.Root.Edges[3].Edges[1].Val)

	serialize(trie, os.Args[2])
	deserialize(os.Args[2])
}