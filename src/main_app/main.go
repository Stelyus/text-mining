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
	"sort"
)

type triple struct {
	Word string
	Freq int
	Distance int
}

func main() {

	a := readInput()
	path := os.Args[1]

	trie := app.Deserialize(path)
	s:= a[0]
	str := strings.Split(s, " ")
	d,_ := strconv.Atoi(str[1])
	res := app.GetDistance(trie, str[2], d)

	sort.Slice(res, func(i, j int) bool {
		if res[i].Distance == res[j].Distance {
			return res[i].Freq > res[j].Freq
		}
		return res[i].Distance < res[j].Distance
	})

	fmt.Printf("[")
	for _,v := range res{
		fmt.Printf("{\"word\":\"%s\",\"freq\":%d,\"distance\":%d},", v.Word, v.Freq, v.Distance)
	}
	fmt.Printf("]\n")

	//fmt.Println(len(res))
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