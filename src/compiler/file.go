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

// Read a file and return the content
func ReadFile(path string) string {
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

func AddWordToTrie(text *string, root *radix.Tree) *radix.Tree {
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

func writeNumberOfEdges(f *os.File, root *radix.Tree) {
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
func writeEdgesSize(f *os.File, root *radix.Tree) {
	bs1 := make([]byte, len(root.Root.Edges) * 4)
    f.Write(bs1)
}

/*
	Serialize the tree into a dict.bin
*/

func Serialize(root *radix.Tree, path string) {
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