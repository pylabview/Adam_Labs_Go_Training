package main

import "fmt"

func main() {
	f := Writer
	fmt.Println(f)
}

// What fmt.Print* does
func Print(v any) {
	if s, ok := v.(fmt.Stringer); ok {
		fmt.Println(s.String())
	}

	fmt.Println(v)
}

// String implement fmt.Stringer
func (f Flag) String() string {
	switch f {
	case Reader:
		return "reader"
	case Writer:
		return "writer"
	case Admin:
		return "admin"
	}

	return fmt.Sprintf("<Flag %d>", f)
}

const (
	Reader Flag = iota + 1
	Writer
	Admin
)

type Flag byte
