package main

import (
	"strings"
	"strconv"
	// "fmt"
	"io/ioutil"
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
	Split a text into an array of (word (string), freq (int))

	Ex:
	
	diieu	2337
	diieux	327

	[(diieu, 2337), (diieux, 327)]
*/

func parseFileToArray(text string) []wordFreq {
	arr := strings.Split(text, "\n")
	var arrayWordFreq []wordFreq

	for i := 0; i < len(arr); i++ {
		fields := strings.Fields(arr[i])
		if len(fields) < 1 {
			continue
		}

		freq, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		arrayWordFreq = append(arrayWordFreq, wordFreq {word: fields[0], freq: freq})
	}
	
	return arrayWordFreq
}

func addWordToTrie(text string, root *Tree) *Tree {
	var first int = 0

	for first != -1 {
		last := strings.IndexByte(text[first:], 10) + first

		if (last == first - 1) {
			break
		}

		fields := strings.Fields(text[first:last])
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