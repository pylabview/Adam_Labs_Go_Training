package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	data, err := os.ReadFile("yankee.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var ms runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("before:", ms.HeapAlloc)

	author, err := findAuthor(data)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	//data = nil
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after :", ms.HeapAlloc)
	//	runtime.KeepAlive(data)

	fmt.Println(string(author))
}

/* Example:
Title: A Connecticut Yankee in King Arthur’s Court

Author: Mark Twain

Release Date: October, 1993 [eBook #86]
*/

func findAuthor(text []byte) ([]byte, error) {
	author := []byte("Author: ")
	i := bytes.Index(text, author)
	if i == -1 {
		return nil, fmt.Errorf("can't find author")
	}

	i += len(author) // Skip "Author: "
	j := bytes.IndexByte(text[i:], '\n')
	if j == -1 {
		return nil, fmt.Errorf("can't find author")
	}

	// return text[i : i+j], nil
	out := make([]byte, j)
	copy(out, text[i:])
	return out, nil
}

// #darkarts in gophers slack
// The compiler does a stack slot liveness analysis.

// go build -gcflags='-alive'
