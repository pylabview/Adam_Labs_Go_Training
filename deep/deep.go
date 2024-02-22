package deep

import "fmt"

// Go does GC shape stenciling code generation
// Two types have the same GC shape if their underlying memory representation is the same
// All pointer in are in the same GC shape

type Ordered interface {
	OrderedNum | string
}

func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("Max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

// [...] is a "type constraint"
type OrderedNum interface {
	~int | ~float64
}

func Relu[T OrderedNum](n T) T {
	if n < 0 {
		return 0
	}

	return n
}

/*
func ReluInt(n int) int {
	if n < 0 {
		return 0
	}

	return n
}

func ReluFloat64(n float64) float64 {
	if n < 0 {
		return 0
	}

	return n
}
*/
