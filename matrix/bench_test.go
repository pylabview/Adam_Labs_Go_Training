package matrix

import (
	"testing"
)

const (
	rows = 1400
	cols = 1050
	/*
		rows = 20
		cols = 70
	*/
)

var (
	mat      = New(rows, cols)
	n        = rows * cols
	expected = (n - 1) * n / 2 // sum of algebraic series
)

func init() {
	for r := range mat.Rows() {
		for c := 0; c < mat.Cols(); c++ {
			mat[r][c] = (r * cols) + c
		}
	}
}

func BenchmarkSumRows(b *testing.B) {
	for range b.N {
		sum := mat.SumRows()
		if sum != expected {
			b.Fatal(sum)
		}
	}
}

func BenchmarkSumCols(b *testing.B) {
	for range b.N {
		sum := mat.SumCols()
		if sum != expected {
			b.Fatal(sum)
		}
	}
}

// go test -bench .
