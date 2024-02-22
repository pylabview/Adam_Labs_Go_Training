package stack

import "fmt"

func ExampleStack() {
	s := New[int]()
	for i := 0; i < 3; i++ {
		s.Push(i)
	}
	fmt.Println("len:", s.Len())
	s.Each(func(n int) bool {
		fmt.Println("each:", n)
		return n > 1
	})
	for s.Len() > 0 {
		v, err := s.Pop()
		fmt.Println("pop:", v, err)
	}
	v, err := s.Pop()
	fmt.Println("pop (empty):", v, err)

	// Output:
	// len: 3
	// each: 2
	// each: 1
	// pop: 2 <nil>
	// pop: 1 <nil>
	// pop: 0 <nil>
	// pop (empty): 0 empty stack
}
