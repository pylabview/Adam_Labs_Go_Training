package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/*
		var mu sync.Mutex // tip: place mutex declaration above variables it guards
		// sync.Mutex can't be copied - always pass a pointer
		counter := 0
	*/
	counter := int64(0)

	var wg sync.WaitGroup
	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1_000; i++ {
				atomic.AddInt64(&counter, 1) // If you reach out for sync/atomic - reconsider
				// If you need metrics: expvar, otel and friends
				/*
					mu.Lock()
					{
						counter++
					}
					mu.Unlock() // Don't forget to Unlock (deadlock) - use defer
					// defer works at function level, this is a loop
				*/
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// go run -race .
// Also build & test
// Common practice: go test -race

// rule 1: If the race detector says there is a race - you have one
// rule 2: If it didn't find a race - you probably have one :)
