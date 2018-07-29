package main

import (
	"fmt"
	"strconv"
	"fmt"
	"bytes"
	"encoding/gob"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	// path := "../ressources/words.txt"
	// wordFreqArr := parseFileToArray(readFile(path))
	// constructTrie(wordFreqArr)

	serialize()
}

type S struct {
	Field1 string
	Field2 int
}

func serialize() {
	s1 := &S{
	    Field1: "Hello Gob",
	    Field2: 999,
	}
	fmt.Println("Original value:", s1)
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(s1)
	if err != nil {
	    fmt.Println("Encode:", err)
	    return
	}

	s2 := &S{}
	err = gob.NewDecoder(buf).Decode(s2)
	if err != nil {
	    fmt.Println("Decode:", err)
	    return
	}

	fmt.Println("New value:", s2)


}

func constructTrie(wordFreqArr []wordFreq ) {
	trie := NewTrie()
	for i := 0; i < len(wordFreqArr); i++ {
		trie.Add(wordFreqArr[i].word, wordFreqArr[i].freq)
 	}
 	node := trie.Root()
	printTrie(node, 0)

}

func printTrie(n *TrieNode, offset int){
	if n == nil {
		return
	}
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	if n.freq != -1 {
		fmt.Print(strconv.QuoteRune(n.val), n.freq)
	} else {
		fmt.Print(strconv.QuoteRune(n.val))
	}


	for _, v := range(n.children){
		printTrie(v, offset + 2)
	}
	fmt.Println()

}