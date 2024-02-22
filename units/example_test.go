package units

import "fmt"

func ExampleValue() {
	v1 := Value{10, Centimeter}
	v2 := Value{3, Inch}

	/* Value semantics
	 */
	v3 := v1.Add(v2)
	fmt.Println(v3)

	/* Pointer semantics
	v1.Add(v2)
	fmt.Println(v1)
	*/

	// Output:
	// {17.62 Centimeter}
}
