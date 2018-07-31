package main

import (
	"os"
	"strings"
	"bufio"
	"app"
	"strconv"
	"fmt"
)

func main() {

	a := readInput()
	path := os.Args[1]

	trie := app.Deserialize(path)

	for _, s := range a {
		str := strings.Split(s, " ")
		d, _ := strconv.Atoi(str[1])

		res := app.GetDistance(trie, str[2], d)
		fmt.Printf("%s\n", app.FormatResult(res))
	}
}



// readInput read input, parse with delimiter '\n' and return lines in an array
func readInput() []string {
	rd := bufio.NewReader(os.Stdin)
	
	var input = []string {}

	s, err := rd.ReadString('\n')
	for err == nil {		
		s = strings.Trim(s, " \n")
		input = append(input, s)
		s, err = rd.ReadString('\n')
	}

	return input
}