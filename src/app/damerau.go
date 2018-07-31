package app

import (
	"radix"
	"sort"
	"math"
	"strconv"
)


// Structure that is use to stock word, frequency and distance
type WordInfo struct {
	Word string
	Freq int
	Distance int
}

// min of two integers
// take two int as imput and return an int
func min(a int, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

// return minimum value in a array
// take an array of int as input, return an int
func minArray(arr []int) int{
	m := math.MaxInt8

	for _, e := range arr {
		if e < m {
			m = e
		}
	}
	return m
}

// create an array of int from min to max
// take two int as input and return an int array
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Call the recursive search on the radix tree and sort the result
// take a radix tree, the word you are looking for and the distance as input
// return an array of WordInfo which contains all the words and their frequency that are at a certain distance
// using damerau-leveistein algorithm

func GetDistance(n *radix.Tree, word string, distance int) []WordInfo {
	var res = []WordInfo {}

	currentRow := makeRange(0, len(word))
	for _, f := range n.Root.Edges{
		searchRecursive(f, word, f.Prefix, currentRow, distance, &res)
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Distance == res[j].Distance {
			if res[i].Freq == res[j].Freq {
				return res[i].Word < res[j].Word
			}
			return res[i].Freq > res[j].Freq
		}
		return res[i].Distance < res[j].Distance
	})

	return res
}

// Recursively search on the radix tree and create a dynamic damerau-leveinsteing distance matrix
// take a node, the word, the currentword that correspond to all the prefix of our branch concatenate, the previous
// row of our matrix and the distance as input
func searchRecursive(node *radix.Node, word string, currentWord string, previousRow []int, maxCost int, res *[]WordInfo){
	// initialize the currentRow with empty array
	columns := len(word) + 1
	currentRow := make([]int, 0)

	// iterate through every letter of the current node
	for t := range node.Prefix {

		// initialize variable
		letter := node.Prefix[t]
		currentRow = make([]int, 0)
		insertCost := 0
		deleteCost := 0
		replaceCost := 0

		// get the first value of the previous row to get the min distance from previous iteration
		currentRow = append(currentRow, previousRow[0]+1)

		// iterate over the length of the word we are looking for
		for column := 1; column < columns; column++ {

			// update the cost depending on the damerau leveinstein matrix
			insertCost = currentRow[column-1] + 1
			deleteCost = previousRow[column] + 1
			subsCost := 0

			// substitution cost
			if word[column-1] != letter {
				subsCost = 1
			}

			replaceCost = previousRow[column - 1] + subsCost

			// if possible try to do a transposition and change the distance accordingly
			d := min(min(insertCost, deleteCost), replaceCost)
			if column > 1 {
				if len(currentWord) > column && len(word) > column && len(currentRow) > column {
					//fmt.Println(len(currentWord), len(word), column)
					if (word[column] == currentWord[column-1]) && (word[column-1] == currentWord[column]) {
						d = min(previousRow[column - 2] + subsCost, currentRow[column])
					}
				}
			}
			currentRow = append(currentRow, d)
		}

		previousRow = currentRow
	}

	// if distance is inferior to maxCost and we are on a leaf we can add the word to our struct
	if currentRow[len(currentRow) - 1] <= maxCost && node.IsLeaf() {
		*res = append(*res, WordInfo{node.Key, int(node.Val), (currentRow[len(currentRow)-1])})
	}

	// search for the minimum value in the currentRow
	m := minArray(currentRow)

	// if the minimum value is inferior to the maxCost, continue the recursion
	// else return that stop the recursion
	if m <= maxCost {
		for _,f := range node.Edges {
			searchRecursive(f, word, currentWord + f.Prefix, currentRow, maxCost, res)
		}
	}
	return
}

// Format the result of the query
func FormatResult(res []WordInfo) string{
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

	return str
}
