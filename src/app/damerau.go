package app

import (
	"radix"
)


/*
		Commente ton code comment j'ai fait dans file.go (meme dossier)
		Quand tu commentes les fonctions commencant par une majuscule, ils doivent commencer par le meme nom que la fonction

		Il faudra aussi commenter un peu a l'interieur de la fonction
*/
type triple struct {
	Word string
	Freq int
	Distance int
}


var out = make(map[string]rune)
var res []triple
var distances = make([]int, 0)

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

func GetDistance(n *radix.Tree, word string, distance int) []triple{
	currentRow := makeRange(0, len(word))
	for _, f := range n.Root.Edges{
		searchRecursive(f, word, f.Prefix, currentRow, distance)
	}
	return res
}

func searchRecursive(node *radix.Node, word string, currentWord string, previousRow []int, maxCost int){
	columns := len(word) + 1
	//fmt.Println(currentWord)
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
			subsCost := 0
			if word[column-1] != letter {
				subsCost = 1
			}

			replaceCost = previousRow[column - 1] + subsCost

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

	if currentRow[len(currentRow) - 1] <= maxCost && node.IsLeaf() {
		//out[node.Key] = node.Val
		res = append(res, triple{node.Key, int(node.Val), (currentRow[len(currentRow)-1])})
		//distances = append(distances, (currentRow[len(currentRow)-1]))
	}

	m := 1000000
	for _, e := range currentRow {
		if e < m {
			m = e
		}
	}
	if m <= maxCost {
		for _,f := range node.Edges {
			searchRecursive(f, word, currentWord + f.Prefix, currentRow, maxCost)
		}
	}
	return
}