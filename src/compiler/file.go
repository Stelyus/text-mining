// Package compiler is used to create TextMiningCompiler.
// It's main function is Serialize that writes a radix.Tree into a file.
package compiler

import (
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/gob"
	"encoding/binary"
    "os"
	"bytes"

	"radix"
)

// ReadFile reads a file and return the content
func ReadFile(path string) string {
	b, err := ioutil.ReadFile(path)

    if err != nil {
        panic(err)
    }

    return string(b)
}



// AddWordToTrie parse text and create a Radix tree.
// The text has a specific format (word 	frequency).
// Return the radix tree created
func AddWordToTrie(text *string, root *radix.Tree) *radix.Tree {
	// We load the text into memory and in order to save memory we move two pointer
	// The first one is "first" which is the beginning of the file and the other one is "last"
	var first int = 0

	for first < len(*text) {
		// Find the first \n which will be "last" position
		last := strings.IndexByte((*text)[first:], 10) + first

		// If there is no \n then we are at our last line
		if (last == first - 1) {
			last = len(*text)						
		}

		// Split " "
		fields := strings.Fields((*text)[first:last])
		if len(fields) < 1 {
			continue
		}

		// Convert frequency into number
		freq, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		// Insert the line parsed into the trie
		root.Insert(fields[0], rune(freq))

		// Update first
		first = last + 1
	}

	return root
}



//	writeNumberOfEdges is used to write at the beginning of the file two bytes that represents the number of the root's edge
func writeNumberOfEdges(f *os.File, root *radix.Tree) {
	bs := make([]byte, 2)
    binary.LittleEndian.PutUint16(bs, uint16(len(root.Root.Edges)))
    f.Write(bs)
}



// Write the size of the encoded node for each root's edge. Each size is encoded with 4 bytes.
// There are initialized with 00 00 00 00, it is replaced at the end of serialize().
// Ex: 30 00 00 00 10 00 00 00.
// That means the first child of the root is encoded with 48 bytes, the next one with 16 bytes.
func writeEdgesSize(f *os.File, root *radix.Tree) {
	bs1 := make([]byte, len(root.Root.Edges) * 4)
    f.Write(bs1)
}


// Serialize serialize the root and write it into path.
// As the builtin serializtion taking too much space, we decided to do it little by little.
// First the file is created. The first two bytes is used to know the number of root's edge (=child).
// The next 4 bytes represents the size of the encoded root's first child in the file, we do this n times with n the number of root's child
// It is initialized with 00 00 00 00 (4 bytes) n times because we don't know yet the size of root's children.
// Then we transform each child into byte and write it to the file. We also save its size.
// Once finished, we go to the beginning of the file with an offset of 2 bytes and replace the initialization with the real
// size of each children. 
func Serialize(root *radix.Tree, path string) {
	var readerBuf bytes.Buffer;
	encoder := gob.NewEncoder(&readerBuf)

	os.Remove(path)
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	// Write the number of root's edge and the size of the encoding for each one of the root's edge.
	writeNumberOfEdges(f, root)
	writeEdgesSize(f, root)




	// encoding each root's edge
	sizePerEdges := make([]int, 0)
	
	for i := 0; i < len(root.Root.Edges); i++ {
		
		err := encoder.Encode(root.Root.Edges[i])

		if err != nil {
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
}