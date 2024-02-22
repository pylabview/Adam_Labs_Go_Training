package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	mu      sync.Mutex
	counter int64
)

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mu.Lock()
		{
			counter++
		}
		mu.Unlock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}
}
