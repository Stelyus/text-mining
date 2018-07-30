package main

import (
	"os"
	"strconv"
	"fmt"
	"encoding/gob"
	"bytes"
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
	// fmt.Println(trie.Root.Edges[3].Edges[1].Prefix)
	// fmt.Println(trie.Root.Edges[3].Edges[1].Val)

	// serialize(trie, os.Args[2])
}

func serialize(node *Tree, path string) {

	var readerBuf bytes.Buffer;
	encoder := gob.NewEncoder(&readerBuf)

	for i := 0; i < len(node.Root.Edges); i++ {
		err := encoder.Encode(node.Root.Edges[i])

		if err != nil {
			panic(err)
			return
		}

		readerBuf.Reset()
	}
}