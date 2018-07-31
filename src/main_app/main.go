package main

import (
	"os"
	// "strconv"
	// "fmt"
	"strings"
	"bufio"
	"app"
	"fmt"
	"strconv"
)

func main() {

	a := readInput()
	path := os.Args[1]

	trie := app.Deserialize(path)

	s:= a[0]
	str := strings.Split(s, " ")
	d,_ := strconv.Atoi(str[1])

	res := app.GetDistance(trie, str[2], d)
	printResult(res)
}


func printResult(res []app.Triple){
	fmt.Printf("[")
	for _,v := range res{
		fmt.Printf("{\"word\":\"%s\",\"freq\":%d,\"distance\":%d},", v.Word, v.Freq, v.Distance)
	}
	fmt.Printf("]\n")
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