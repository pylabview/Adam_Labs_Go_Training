package main

import (
	"testing"
	"time"
)

var (
	ae = AuthError{
		Login:  "elliot",
		Time:   time.Now().UTC(),
		Reason: "bad password",
	}
	e error = &ae
	s string
)

func BenchmarkMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = ae.Error() // assign to s so compile won't optimize away
	}
}

func BenchmarkIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = e.Error() // assign to s so compile won't optimize away
	}
}
