package main

import (
	"os"
	// "strconv"
	// "fmt"
	"strings"
	"bufio"

	"app"
)

func main() {

	readInput()

	path := os.Args[1]
	// distance, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	panic(err)
	// }

	// word := os.Args[3]

	app.Deserialize(path)
	// out := app.Testalgo(trie, word, distance)

	// for k,v := range out{
	// 	fmt.Println("{\"word\": \"", k, "\"freq\":", v, "\"distance\":", 2)
	// }
	// fmt.Println(len(out))	
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