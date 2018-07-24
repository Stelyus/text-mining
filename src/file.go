package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }

    return string(b) // convert content to a 'string'
}