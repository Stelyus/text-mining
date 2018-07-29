package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"encoding/gob"
    "os"
)

type wordFreq struct {
	word string
	freq int
}


// Read a file and return the content
func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

    if err != nil {
        panic(err)
    }

    return string(b) // convert content to a 'string'
}


/*
	Parse text and create a Radix tree
	Return the radix tree created
*/

func addWordToTrie(text *string, root *Tree) *Tree {
	var first int = 0

	for first < len(*text) {
		last := strings.IndexByte((*text)[first:], 10) + first

		if (last == first - 1) {
			last = len(*text)						
		}

		fields := strings.Fields((*text)[first:last])
		if len(fields) < 1 {
			continue
		}

		freq, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		root.Insert(fields[0], rune(freq))

		first = last + 1
	}

	return root
}


/*
	Serialize the tree into a dict.bin
*/

func serialize(node *Tree, path string) {
	encodeFile, _ := os.Create(path)

	err := gob.NewEncoder(encodeFile).Encode(*node)

	if err != nil {
	    fmt.Println("Encode:", err)
	    return
	}

	encodeFile.Close()


	// Decode
	// newNode := &Tree{}

	// decodeFile, _ := os.Open("dict_berthang.bin")
	// err = gob.NewDecoder(decodeFile).Decode(newNode)
	// if err != nil {
	//     fmt.Println("Decode:", err)
	//     return
	// }

	// fmt.Println("New value:", newNode)
	// fmt.Println(newNode.Root.Edges[1].Edges[1].Leaf.Key)


}