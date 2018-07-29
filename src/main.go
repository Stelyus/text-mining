package main

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := "../ressources/words.txt"
	//wordFreqArr := parseFileToArray(readFile(path))
	trie := NewRadix()
	str := readFile(path)
	trie = addWordToTrie(&str, trie)
	//constructTrie(wordFreqArr)

	// serialize()
}

//
//func serialize(node *TrieNode) {
//	fmt.Println("Original value:", *node)
//	buf := new(bytes.Buffer)
//	err := gob.NewEncoder(buf).Encode(*node)
//	if err != nil {
//	    fmt.Println("Encode:", err)
//	    return
//	}
//
//	newNode := &TrieNode{}
//	err = gob.NewDecoder(buf).Decode(newNode)
//	if err != nil {
//	    fmt.Println("Decode:", err)
//	    return
//	}
//
//	fmt.Println("New value:", newNode)
//
//
//}

func constructTrie(wordFreqArr []wordFreq ) *Tree {
	trie := NewRadix()
	for i := 0; i < len(wordFreqArr); i++ {
		trie.Insert(wordFreqArr[i].word, rune(wordFreqArr[i].freq))
 	}
	//printTrie(trie, 0)

	return trie
}