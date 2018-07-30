package main

import ("os")

func main() {
	// Si lancer depuis autre que src, le path n'est pas correct
	// distance, _ := strconv.Atoi(os.Args[3])
	// word := os.Args[4]

	// out := getwords(trie, word, distance)
	// out := getwords(trie, word, distance)
	//fmt.Println(DamerauLevenshtein(word, "abilo"), distance)
	//fmt.Println(out)

	// for k,v := range out{
	// 	fmt.Println("{\"word\": \"", k, "\"freq\":", v, "\"distance\":", 2)
	// }
	// fmt.Println(len(out))

	deserialize(os.Args[1])
}