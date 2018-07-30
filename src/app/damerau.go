package main

import (
	"radix"
)


var out = make(map[string]rune)
// min of two integers
func min(a int, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func testalgo(n *radix.Tree, word string, distance int) map[string]rune{
	currentRow := makeRange(0, len(word))
	for _, f := range n.Root.Edges{
		searchRecursive(f, word, currentRow, distance)
	}
	return out
}

func searchRecursive(node *radix.Node, word string, previousRow []int, maxCost int){
	columns := len(word) + 1

	currentRow := make([]int, 0)

	for t := range node.Prefix {

		letter := node.Prefix[t]
		currentRow = make([]int, 0)
		currentRow = append(currentRow, previousRow[0]+1)

		insertCost := 0
		deleteCost := 0
		replaceCost := 0

		for column := 1; column < columns; column++ {

			insertCost = currentRow[column-1] + 1
			deleteCost = previousRow[column] + 1

			if word[column-1] != letter {
				replaceCost = previousRow[ column-1 ] + 1
			} else {
				replaceCost = previousRow[ column-1 ]
			}
			currentRow = append(currentRow, min(min(insertCost, deleteCost), replaceCost))
		}

		previousRow = currentRow
	}

	if currentRow[len(currentRow) - 1] <= maxCost && node.IsLeaf() {
		out[node.Key] = rune(currentRow[len(currentRow)-1])
	}

	m := 1000000
	for _, e := range currentRow {
		if e < m {
			m = e
		}
	}
	if m <= maxCost {
		for _,f := range node.Edges {
			searchRecursive(f, word, currentRow, maxCost)
		}
	}
	return
}