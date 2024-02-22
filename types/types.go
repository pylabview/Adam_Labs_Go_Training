package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	// value semantics for built-in types
	s := "hello"
	u := strings.ToUpper(s)
	fmt.Println(u)

	// type run = int32
	data := []rune{'a', 'b', 'c'}
	fmt.Printf("start: %p\n", &data[0])
	fmt.Printf("end  : %p\n", &data[2])

	cart := []string{"apple", "lemon", "banana"}
	cart = append(cart, "bread")
	fmt.Println("cart:", cart, "len:", len(cart), "cap:", cap(cart))
	// fruit := cart[:3]
	fruit := cart[:3:3]
	fmt.Println("fruit:", fruit, "len:", len(fruit), "cap:", cap(fruit))
	fruit = append(fruit, "pear")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart)

	const size = 2_000
	var nums []int
	// If you know the size in advance - pre allocate
	// nums := make([]int, 0, size)
	for i := range size {
		nums = appendInt(nums, i)
	}
	fmt.Println(nums[:10])

	values := []float64{3, 1, 2}
	fmt.Println(median(values))
	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values))
	fmt.Println(values)

	var s1 []rune
	s2 := []rune{}
	fmt.Println(reflect.DeepEqual(s1, s2))
	// Prefer `if len(s) == 0` to `if s == nil`
}

/*
median

- sort values
- if odd number of values, return middle
- return average of middles
*/
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// Local copy in order not to change values
	vs := make([]float64, len(values))
	copy(vs, values)
	/*
		values = append(values, 0)
		vs := values
	*/

	sort.Float64s(vs)
	i := len(vs) / 2
	if len(vs)%2 == 1 { // odd
		return vs[i], nil
	}

	m := (vs[i-1] + vs[i]) / 2
	return m, nil
}

func appendInt(s []int, v int) []int {
	if len(s) == cap(s) { // underlying array is full: re-allocate & copy
		// Go has a different grow algorithms which is undefined
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		s2 := make([]int, size)
		copy(s2, s)
		s = s2[:len(s)]
	}

	s = s[:len(s)+1]
	s[len(s)-1] = v
	return s
}

/*
cart
	[apple lemon banana bread "" ""]
	array = &cart[0]
	len = 4 (how much we look at)
	cap = 6 (size of underlying array from start)
*/

/* Three class of types
built-in/scalar: numbers, strings, byte, runes...
	- Value semantics
	- Exception: marshaling
internal: slices, maps, channels
	- channels & maps are pointers already - pass by value
	- slices no a structs - pass by value
user defined:
	- If the underlying type is one of above - do the same
		e.g. type Number int
	- structs
		- Try working with value semantics
		- Use pointer if have a mutating method
		- Use pointer if you have a lock (sync.Mutex)
		- For fields, use the same rules (mostly by value)
*/
