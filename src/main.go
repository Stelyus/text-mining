package main

import (
	"strconv"
	"bytes"
	"fmt"
	"encoding/gob"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := "../ressources/test.txt"
	wordFreqArr := parseFileToArray(readFile(path))
	constructTrie(wordFreqArr)

	// serialize()
}


func serialize(node *TrieNode) {
	fmt.Println("Original value:", *node)
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(*node)
	if err != nil {
	    fmt.Println("Encode:", err)
	    return
	}

	newNode := &TrieNode{}
	err = gob.NewDecoder(buf).Decode(newNode)
	if err != nil {
	    fmt.Println("Decode:", err)
	    return
	}

	fmt.Println("New value:", newNode)


}

func constructTrie(wordFreqArr []wordFreq ) *TrieNode {
	trie := NewTrie()
	for i := 0; i < len(wordFreqArr); i++ {
		trie.Add(wordFreqArr[i].word, wordFreqArr[i].freq)
 	}
	printTrie(trie, 0)

	return trie
}

func printTrie(n *TrieNode, offset int){
	if n == nil {
		return
	}
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	if n.Freq != -1 {
		fmt.Print(strconv.QuoteRune(n.Val), n.Freq)
	} else {
		fmt.Print(strconv.QuoteRune(n.Val))
	}

	for _, v := range(n.Children){
		printTrie(v, offset + 2)
	}
	fmt.Println()

}