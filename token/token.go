package main

import (
	"fmt"
	"sync"
	"time"
)

type Token struct {
	mu    sync.RWMutex
	token string

	n int
}

func NewToken() *Token {
	var t Token
	t.token = t.refreshToken()

	go func() {
		for {
			time.Sleep(time.Second)
			t.mu.Lock() // writer lock
			{
				fmt.Println("getting new token")
				t.token = t.refreshToken()
				fmt.Printf("new token: %s\n", t.token)
			}
			t.mu.Unlock()
		}
	}()

	return &t
}

func (t *Token) Get() string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.token
}

func (t *Token) refreshToken() string {
	time.Sleep(100 * time.Millisecond)
	t.n++
	return fmt.Sprintf("Token %d", t.n)
}

func main() {
	t := NewToken()
	for range 10 {
		fmt.Println("getting token")
		start := time.Now()
		tok := t.Get()
		fmt.Printf("got %s in %v\n", tok, time.Since(start))
		time.Sleep(200 * time.Millisecond)
	}
}
