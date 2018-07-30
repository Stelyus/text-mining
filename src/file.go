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
	Serialize the tree into a dict.bin
*/

func serialize(root *Tree, path string) {
	var readerBuf bytes.Buffer;
	encoder := gob.NewEncoder(&readerBuf)

	os.Remove(path)
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)

	// WRITE TWO BYTES
	bs := make([]byte, 2)
    binary.LittleEndian.PutUint16(bs, uint16(len(root.Root.Edges)))
    f.Write(bs)

    // WRITE 4 * LEN EDGES BYTES
	bs1 := make([]byte, len(root.Root.Edges) * 4)
    f.Write(bs1)



	sizePerEdges := make([]int, 0)
	
	for i := 0; i < len(root.Root.Edges); i++ {
		
		err := encoder.Encode(root.Root.Edges[i])

		if err != nil {
			f.Close()
			panic(err)
		}

		sizePerEdges = append(sizePerEdges, readerBuf.Len())
		f.Write(readerBuf.Bytes())
		readerBuf.Reset()
	}


    var i int64

    for i = 0; i < int64(len(sizePerEdges)); i++ {
		bs := make([]byte, 4)
    	binary.LittleEndian.PutUint32(bs, uint32(sizePerEdges[i]))
    	f.WriteAt(bs, (i * 4) + 2)
    }


	f.Close()
}

func deserialize(path string) {

	trie := NewRadix()


	/*
		Calcul du nombre d'edge (2 premier bytes)
	*/

	numberEdgesByte := make([]byte, 2)
	f, err := os.Open(path)
	check(err)

	_, err1 := f.Read(numberEdgesByte)
    check(err1)
    numberEdges := binary.LittleEndian.Uint16(numberEdgesByte)


    /*
		Pour chaque edge, calcul la taille de chaque node en byte
    */

    sizePerEdges := make([]uint32, 0)

    var i uint16
    for i = 0; i < numberEdges; i++ {
    	bs := make([]byte, 4)
    	f.Read(bs)
    	sizePerEdges = append(sizePerEdges, binary.LittleEndian.Uint32(bs))
    }


    /*
		Creation de l'arbre
    */

	var readerBuf bytes.Buffer
	decoder := gob.NewDecoder(&readerBuf)


    for i := 0; i < len(sizePerEdges); i++ {
    	readerBuf.Grow(int(sizePerEdges[i]))
    	bs := make([]byte, sizePerEdges[i])

    	node := &node{}
    	f.Read(bs)
    	readerBuf.Write(bs)
    	decoder.Decode(node)
    	readerBuf.Reset()
    	trie.Root.addEdge(node)
    }

    fmt.Println("Deserialization done")
    f.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}