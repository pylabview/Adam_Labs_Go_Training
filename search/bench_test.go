package main

import "testing"

func BenchmarkQuery(b *testing.B) {
	db := NewDB()
	for i := 0; i < b.N; i++ {
		items := db.Search("president")
		if len(items) != 46 {
			b.Fatal(len(items))
		}
	}
}

/* RPS
 */
