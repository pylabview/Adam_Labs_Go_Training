package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var nReaders int64

func main() {
	db := DB{
		data: make(map[string]int),
	}
	const key = "goroutines.iter"
	var wg sync.WaitGroup

	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range 10_000 {
				db.Set(key, n)
			}
		}()
	}

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10_000 {
				db.Get(key)
			}
		}()
	}

	wg.Wait()
}

func (db *DB) Get(key string) int {
	db.mu.RLock() // Reader lock
	defer db.mu.RUnlock()

	defer atomic.AddInt64(&nReaders, -1)

	n := atomic.AddInt64(&nReaders, 1)
	fmt.Printf("get: nReaders = %d\n", n)

	randSleep()
	return db.data[key]
}

func (db *DB) Set(key string, value int) {
	db.mu.Lock() // Writer lock
	defer db.mu.Unlock()

	n := atomic.AddInt64(&nReaders, 0)
	fmt.Println("set: nReaders =", n)
	fmt.Printf("\033[31mset: nReader = %d\033[39m\n", n)

	randSleep()
	db.data[key] = value
}

type DB struct {
	mu   sync.RWMutex
	data map[string]int
}

func randSleep() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}
