package main

import (
	"fmt"
)

func main() {
	text := "make the zero value useful"
	freqs := make(map[rune]int) // char -> count
	for _, r := range text {
		freqs[r]++ // read, modify, write
		// n = freqs[r]
		// n++
		// freqs[r] = n
	}

	for r, c := range freqs {
		fmt.Printf("%c: %d\n", r, c)
	}

	c := freqs['b']
	fmt.Println("b count:", c) // zero vs missing
	c, ok := freqs['b']        // "comma ok" (maps, type assertions, channels)
	if !ok {
		fmt.Println("b not in freq")
	} else {
		fmt.Println("b count:", c)
	}
}
