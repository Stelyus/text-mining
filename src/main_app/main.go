package main

import (
	"os"
	"strconv"
	"fmt"

	"app"
)

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	path := os.Args[1]
	distance, _ := strconv.Atoi(os.Args[2])
	word := os.Args[3]
	// trie := NewRadix()
	// str := readFile(path)
	// trie = addWordToTrie(&str, trie)

	trie := app.Deserialize(path)
	out := app.Testalgo(trie, word, distance)

	for k,v := range out{
		fmt.Println("{\"word\": \"", k, "\"freq\":", v, "\"distance\":", 2)
	}
	fmt.Println(len(out))

	// serialize(trie, os.Args[2])
	
}