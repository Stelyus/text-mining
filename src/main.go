package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Si lancer depuis autre que src, le path n'est pas correct
	path := "../ressources/test.txt"
	wordFreqArr := parseFileToArray(readFile(path))
	constructTrie(wordFreqArr)
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