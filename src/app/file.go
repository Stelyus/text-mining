// Package app is used to create TextMiningApp.
// There is two main features, the deserialization and the research using Damerau–Levenshtein distance
package app

import (
    "radix"
    "encoding/gob"
    "encoding/binary"
    "os"
    "bytes"
)



// Read the two first byte of the file, and return the number in uint16.
// The result represents the number of Root's Edge
func numberRootsEdge(f *os.File) uint16 {
	numberEdgesByte := make([]byte, 2)
	_, err := f.Read(numberEdgesByte)
	check(err)
    return binary.LittleEndian.Uint16(numberEdgesByte)
}


// Read the next four bytes in order to know what is the length of the encoded edges
// Put the size into an array and return it
// Ex: 05 00 00 00    0A 00 00 00
// That means the first 4 bits represents the first edge and the size of the encoded edge 5
// The next 4 represents the second edge and is encoded in 10 bytes.
// After reading numberEdges * 4 bytes, there is the data
// The file use Little Endian notation.
func bytesPerEdge(f *os.File, numberEdges uint16) []uint32 {

	sizePerEdges := make([]uint32, numberEdges)

    var i uint16
    for i = 0; i < numberEdges; i++ {
    	bs := make([]byte, 4)
    	f.Read(bs)
    	sizePerEdges[i] = binary.LittleEndian.Uint32(bs)
    }

    return sizePerEdges
}

// Deserialize take a path in argument and return a radix.Tree
// First it creates the root.
// It reads the first two bytes which is the number of root's child (numberRootsEdge()).
// And then it read the next n * 4 bytes in order to know what is the size of each child bytes size.
// Finally it creates the nodes which will be append with the root.
func Deserialize(path string) *radix.Tree {

		
	// Create the root
	trie := radix.NewRadix()

	f, err := os.Open(path)
    defer f.Close()
	check(err)

	numberEdges := numberRootsEdge(f)
	sizePerEdges := bytesPerEdge(f, numberEdges)

	// Create children root
	var readerBuf bytes.Buffer
	decoder := gob.NewDecoder(&readerBuf)


    for i := 0; i < len(sizePerEdges); i++ {
    	// Initialization, reading sizePerEdges[i] bytes
    	readerBuf.Grow(int(sizePerEdges[i]))
    	bs := make([]byte, sizePerEdges[i])

    	

    	// Reading and transfer it to a buffer
    	f.Read(bs)
    	readerBuf.Write(bs)

    	// Decode it to a node
    	node := &radix.Node{}
    	decoder.Decode(node)

    	// Reset the buffer so that it can grow to sizePerEdges[i+1] bytes
    	readerBuf.Reset()

		// Add it to the root's edge    	
    	trie.Root.AddEdge(node)
    }

    return trie
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}