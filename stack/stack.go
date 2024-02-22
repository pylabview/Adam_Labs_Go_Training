package stack

import "errors"

var ErrEmpty = errors.New("empty stack")

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	head *node[T]
	size int
}

func New[T any]() *Stack[T] {
	s := Stack[T]{}
	return &s
}

func (s *Stack[T]) Len() int {
	return s.size
}

func (s *Stack[T]) Push(val T) {
	s.head = &node[T]{val, s.head}
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var zero T
		return zero, ErrEmpty
	}

	n := s.head
	s.head = n.next
	s.size--

	return n.value, nil
}

func (s *Stack[T]) Each(fn func(T) bool) {
	n := s.head
	for n != nil {
		if !fn(n.value) {
			return
		}
		n = n.next
	}
}

// This won't compile
// func (s *Stack[T])[V any] AddHead(v V) {

// type Cache[K comparable, V any] map[K]V
