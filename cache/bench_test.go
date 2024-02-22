package cache

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

var (
	sc   SliceCache
	mc   = make(MapCache)
	size = 1000
)

func init() {
	for i := range size {
		login := fmt.Sprintf("user-%03d", i)
		u := User{Login: login}
		sc = append(sc, u)
		mc[login] = u
	}

	rand.Shuffle(size, func(i, j int) { sc[i], sc[j] = sc[j], sc[i] })
}

// I have 80 cache hits (need monitoring to know)

func BenchmarkSliceCache(b *testing.B) {
	/*
		b.StopTimer()
		// ...
		b.StartTimer()
		b.ResetTimer()
	*/
	for range b.N {
		for range 2 { // 20% cache miss
			_, ok := sc.Get("no such user")
			if ok {
				b.Fatalf("found non existing user")
			}
		}
		for range 8 { // 80% cache hit
			_, ok := sc.Get("user-001")
			if !ok {
				b.Fatalf("not found")
			}
		}
	}
}

func BenchmarkMapCache(b *testing.B) {
	for range b.N {
		for range 2 { // 20% cache miss
			_, ok := mc.Get("no such user")
			if ok {
				b.Fatalf("found non existing user")
			}
		}
		for range 8 { // 80% cache hit
			_, ok := mc.Get("user-001")
			if !ok {
				b.Fatalf("not found")
			}
		}
	}
}

// go test -run NONE -bench . -count 10
