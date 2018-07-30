package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"encoding/gob"
	"encoding/binary"
    "os"
	"bytes"
)

type wordFreq struct {
	word string
	freq int
}


// Read a file and return the content
func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

    if err != nil {
        panic(err)
    }

    return string(b) // convert content to a 'string'
}


/*
	Parse text and create a Radix tree
	Return the radix tree created
*/

func addWordToTrie(text *string, root *Tree) *Tree {
	var first int = 0

	for first < len(*text) {
		last := strings.IndexByte((*text)[first:], 10) + first

		if (last == first - 1) {
			last = len(*text)						
		}

		fields := strings.Fields((*text)[first:last])
		if len(fields) < 1 {
			continue
		}

		freq, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		root.Insert(fields[0], rune(freq))

		first = last + 1
	}

	return root
}


/*
	Write at the beginning of the file, two bytes that represents the number of the root's edge
*/

func writeNumberOfEdges(f *os.File, root *Tree) {
	bs := make([]byte, 2)
    binary.LittleEndian.PutUint16(bs, uint16(len(root.Root.Edges)))
    f.Write(bs)
}


/*
	Write the size of the encoded node for each root's edge. Each size is encoded with 4 bytes.
	There are initialized with 00 00 00 00, it is replaced at the end of serialize()

	Ex: 30 00 00 00 10 00 00 00
	That means the first child of the root is encoded with 48 bytes, the next one with 16 bytes
*/
func writeEdgesSize(f *os.File, root *Tree) {
	bs1 := make([]byte, len(root.Root.Edges) * 4)
    f.Write(bs1)
}

/*
	Serialize the tree into a dict.bin
*/

func serialize(root *Tree, path string) {
	var readerBuf bytes.Buffer;
	encoder := gob.NewEncoder(&readerBuf)

	os.Remove(path)
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)

	// Write the number of root's edge and the size of the encoding for each one of the root's edge.
	writeNumberOfEdges(f, root)
	writeEdgesSize(f, root)




	// Encoded each root's edge
	sizePerEdges := make([]int, 0)
	
	for i := 0; i < len(root.Root.Edges); i++ {
		
		err := encoder.Encode(root.Root.Edges[i])

		if err != nil {
			f.Close()
			panic(err)
		}

		// Save the size of the encoded node
		sizePerEdges = append(sizePerEdges, readerBuf.Len())

		// Write the encoded node in the file
		f.Write(readerBuf.Bytes())
		readerBuf.Reset()
	}



	// Replace the header with the corresponding encoded size
    var i int64

    for i = 0; i < int64(len(sizePerEdges)); i++ {
		bs := make([]byte, 4)
    	binary.LittleEndian.PutUint32(bs, uint32(sizePerEdges[i]))
    	f.WriteAt(bs, (i * 4) + 2)
    }


	f.Close()
}


/*
	Read the two first byte of the file, and return the number in uint16.
	The result represents the number of Root's Edge
*/

func numberRootsEdge(f *os.File) uint16 {
	numberEdgesByte := make([]byte, 2)
	_, err := f.Read(numberEdgesByte)
	check(err)
    return binary.LittleEndian.Uint16(numberEdgesByte)
}


/*
	Read the next four bytes in order to know what is the length of the encoded edges
	Put the size into an array and return it

	Ex: 05 00 00 00    0A 00 00 00
	That means the first 4 bits represents the first edge and the size of the encoded edge 5
	The next 4 represents the second edge and is encoded in 10 bytes.
	After reading numberEdges * 4 bytes, there is the data
	The file use Little Endian notation.
*/

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


func deserialize(path string) *Tree {

	/*
		Create the root
	*/
	trie := NewRadix()

	f, err := os.Open(path)
	check(err)

	numberEdges := numberRootsEdge(f)
	sizePerEdges := bytesPerEdge(f, numberEdges)

    /*
		Creation de l'arbre
    */

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
    	node := &node{}
    	decoder.Decode(node)

    	// Reset the buffer so that it can grow to sizePerEdges[i+1] bytes
    	readerBuf.Reset()

		// Add it to the root's edge    	
    	trie.Root.addEdge(node)
    }

    fmt.Println("Deserialization done")
    f.Close()

    return trie
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}