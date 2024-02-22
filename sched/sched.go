package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	/*
		for msg := queue {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// ...
			}()
		}
	*/

	const niter = 100
	fmt.Println("Starting")

	go func() {
		for n := 0.0; n < niter; n++ {
			s := Sqrt(n)
			// Blue
			fmt.Printf("\033[32m%.2f->%.2f\033[39m ", n, s)
		}
		wg.Done()
	}()

	go func() {
		for n := 0.0; n < niter; n++ {
			s := Sqrt(n)
			// Red
			fmt.Printf("\033[31m%.2f->%.2f\033[39m ", n, s)
		}
		wg.Done()
	}()

	fmt.Println("Waiting")
	wg.Wait()

	fmt.Println("\nDone")
}

// Sqrt calculate the square root of val using Newton's method
func Sqrt(val float64) float64 {
	if val == 0.0 {
		return 0.0
	}

	guess, epsilon := 1.0, 0.0001
	for i := 0; i < 10_000; i++ {
		if Abs(guess*guess-val) <= epsilon {
			return guess
		}
		guess = (val/guess + guess) / 2.0
	}

	return guess
}

// Abs return the absolute value of val
func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}

	return val
}

// GOMAXPROCS=1  go run .
// GOMAXPROCS=2  go run .
// GOMAXPROCS=1 GODEBUG=schedtrace=1 go run .
