package main

import (
	"fmt"
	"strings"
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

// max of two integers
func maxI(a int, b int) (res int) {
	if a < b {
		res = b
	} else {
		res = a
	}

	return
}

func DamerauLevenshtein(s1 string, s2 string) (distance int) {
	r1 := []rune(s1)
	r2 := []rune(s2)

	// the maximum possible distance
	inf := len(r1) + len(r2)

	// construct the edit-tracking matrix
	matrix := make([][]int, len(r1))
	for i := range matrix {
		matrix[i] = make([]int, len(r2))
	}

	// seen characters
	seenRunes := make(map[rune]int)

	if r1[0] != r2[0] {
		matrix[0][0] = 1
	}

	seenRunes[r1[0]] = 0
	for i := 1; i < len(r1); i++ {
		deleteDist := matrix[i-1][0] + 1
		insertDist := (i+1)*1 + 1
		var matchDist int
		if r1[i] == r2[0] {
			matchDist = i
		} else {
			matchDist = i + 1
		}
		matrix[i][0] = min(min(deleteDist, insertDist), matchDist)
	}

	for j := 1; j < len(r2); j++ {
		deleteDist := (j + 1) * 2
		insertDist := matrix[0][j-1] + 1
		var matchDist int
		if r1[0] == r2[j] {
			matchDist = j
		} else {
			matchDist = j + 1
		}

		matrix[0][j] = min(min(deleteDist, insertDist), matchDist)
	}

	for i := 1; i < len(r1); i++ {
		var maxSrcMatchIndex int
		if r1[i] == r2[0] {
			maxSrcMatchIndex = 0
		} else {
			maxSrcMatchIndex = -1
		}

		for j := 1; j < len(r2); j++ {
			swapIndex, ok := seenRunes[r2[j]]
			jSwap := maxSrcMatchIndex
			deleteDist := matrix[i-1][j] + 1
			insertDist := matrix[i][j-1] + 1
			matchDist := matrix[i-1][j-1]
			if r1[i] != r2[j] {
				matchDist += 1
			} else {
				maxSrcMatchIndex = j
			}

			// for transpositions
			var swapDist int
			if ok && jSwap != -1 {
				iSwap := swapIndex
				var preSwapCost int
				if iSwap == 0 && jSwap == 0 {
					preSwapCost = 0
				} else {
					preSwapCost = matrix[maxI(0, iSwap-1)][maxI(0, jSwap-1)]
				}
				swapDist = i + j + preSwapCost - iSwap - jSwap - 1
			} else {
				swapDist = inf
			}
			matrix[i][j] = min(min(min(deleteDist, insertDist), matchDist), swapDist)
		}
		seenRunes[r1[i]] = i
	}

	return matrix[len(r1)-1][len(r2)-1]
}


func getwords(n *radix.Tree, word string, distance int) map[string]rune{
	//currentRow := makeRange(0, len(word) + 1)

	for _, f := range n.Root.Edges{
		recursives(f, word, distance, f.Prefix)
	}
	
	return out
}

func recursives(n *radix.Node, s string, distance int, check string){
	fmt.Println("--------------------------")
	fmt.Println("recursive call on:", check)
	if n.IsLeaf() {
		fmt.Println("checking leaf", n.Key)
		if DamerauLevenshtein(s, n.Key) <= distance {
			out[n.Key] = n.Val
		}
	}

	for _, f := range n.Edges{
		if !checkDamerau(s, check + f.Prefix, distance, f.IsLeaf()){
			recursives(f, s, distance, check + f.Prefix)
		}
	}
	fmt.Println("==============", check)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

//# This recursive helper is used by the search function above. It assumes that
//# the previousRow has been filled in already.
//func searchRecursive(node *node, letter string, word string, previousRow []int, results []int, maxCost int){
// colums := makeRange(0, len(word) + 1)
//
//}
//currentRow = [ previousRow[0] + 1 ]

//# Build one row for the letter, with a column for each letter in the target
//# word, plus one for the empty string at column 0
//for column in xrange( 1, columns ):
//	insertCost = currentRow[column - 1] + 1
//	deleteCost = previousRow[column] + 1
//
//	if word[column - 1] != letter:
//		replaceCost = previousRow[ column - 1 ] + 1
//	else:
//		replaceCost = previousRow[ column - 1 ]
//
//	currentRow.append( min( insertCost, deleteCost, replaceCost ) )

//# if the last entry in the row indicates the optimal cost is less than the
//# maximum cost, and there is a word in this trie node, then add it.
//if currentRow[-1] <= maxCost and node.word != None:
//	out[node.Key] = currentRow[-1]

//# if any entries in the row are less than the maximum cost, then
//# recursively search each branch of the trie
//if min( currentRow ) <= maxCost:
//	for letter in node.children:
//		searchRecursive( node.children[letter], letter, word, currentRow, results, maxCost )

func checkDamerau(ref string, currentWord string, distance int, isleaf bool) bool{

	//fmt.Println(ref, currentWord, distance)
	// check if length of the current word doesn't go way over distance
	if len(currentWord) - len(ref) > distance {
		return true
	}

	if currentWord != "" {
		d := 0
		// get the length of the current prefix
		mini := min(len(ref), len(currentWord))

		if len(ref) > mini {
			// check the prefix of the ref and the suffix to correctly check for damerau leveisntein distance
			prefix := ref[:mini]
			suffix := ref[1 : mini+1]
			if strings.Contains(currentWord, "nai") {
				fmt.Println("prefixe2:", prefix, suffix, currentWord)

			}
			d = min(DamerauLevenshtein(prefix, currentWord), DamerauLevenshtein(suffix, currentWord))
		} else {
			currentWord = currentWord[:mini]
			d = DamerauLevenshtein(ref, currentWord)
		}

		d = min(d, DamerauLevenshtein(ref, currentWord))
		if strings.Contains(currentWord, "nai") {

		fmt.Println("prefixe:", ref, currentWord)
		fmt.Println("distance is:", d)
		fmt.Println("---")
		}
		return !(d <= distance)
	}
	return false
}
