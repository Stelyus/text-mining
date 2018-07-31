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


func printResult(res []app.Triple) string{
	i := 1
	str := "["
	for _,v := range res{
		if i != len(res) {
			str += "{\"word\":\"" + v.Word + "\",\"freq\":" + strconv.Itoa(v.Freq) + ",\"distance\":"+ strconv.Itoa(v.Distance) + "},"
		}else{
			str += "{\"word\":\"" + v.Word + "\",\"freq\":" + strconv.Itoa(v.Freq) + ",\"distance\":"+ strconv.Itoa(v.Distance) + "}"
		}
		i++
	}
	str += "]"
	fmt.Printf("%s\n", str)
	return str
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