package units

import (
	"fmt"
	"math/big"
)

type Unit struct {
	name string
}

func (u Unit) String() string {
	return u.name
}

var (
	Centimeter = Unit{"Centimeter"}
	Inch       = Unit{"Inch"}
)

// 1 inch = 2.54 centimeters

type Value struct {
	Amount float64
	Unit   Unit
}

type regKey struct {
	from Unit
	to   Unit
}

var registry = map[regKey]func(float64) float64{
	{Centimeter, Inch}: func(v float64) float64 { return v / 2.54 },
	{Inch, Centimeter}: func(v float64) float64 { return v * 2.54 },
}

/* Pointer semantics
func (v *Value) Add(other Value) {
	a := other.Amount
	if v.Unit != other.Unit {
		a = registry[regKey{other.Unit, v.Unit}](a)
	}

	v.Amount += a
}
*/

/* Value semantics
 */
func (v Value) Add(other Value) Value {
	a := other.Amount
	if v.Unit != other.Unit {
		a = registry[regKey{other.Unit, v.Unit}](a)
	}

	return Value{v.Amount + a, v.Unit}
}

/* math/big API uses pointers semantics */
func bigAPI() {
	a := big.NewInt(7)
	b := big.NewInt(3)
	// a += b
	a = a.Add(a, b)
	// a = big.Add(a, b)
	fmt.Println(a)
}
