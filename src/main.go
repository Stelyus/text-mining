package main


import "fmt"

func main() {

	// Si lancer depuis autre que src, le path n'est pas correct
	path := "../ressources/words.txt"
	wordFreqArr := parseFileToArray(readFile(path))
	constructTrie(wordFreqArr)
}


func constructTrie(wordFreqArr []wordFreq ) {
	trie := TrieConstructor();
	for i := 0; i < len(wordFreqArr); i++ {
		trie.Insert(wordFreqArr[i].word);
 	}
 	
 	fmt.Println(trie.Search("n675"))
}