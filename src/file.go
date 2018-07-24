package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
)

type wordFreq struct {
	word string
	freq int
}

func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }

    return string(b) // convert content to a 'string'
}


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
			fmt.Println(err)
			os.Exit(3)
		}

		arrayWordFreq = append(arrayWordFreq, wordFreq {word: fields[0], freq: freq})
	}
	
	return arrayWordFreq
}