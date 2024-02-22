package user_test

import (
	"fmt"

	"ultimate-go/user"
)

func ExampleNew() {
	u := user.New("elliot", "elliot@e-corp.com")
	fmt.Println(u)

	// Output:
	// &{elliot elliot@e-corp.com}
}
