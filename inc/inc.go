package main

import "fmt"

func main() {
	n := 41
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
	inc(&n)
	fmt.Printf("main: n=%d, addr=%p\n", n, &n)
}

func inc(n *int) {
	(*n)++ // read, modify, write
	fmt.Printf("inc : n=%d, addr=%p\n", *n, n)
}
