package main

import "fmt"

type Account struct {
	Login   string
	Type    string
	Balance int // ¢
}

func main() {
	bank := []Account{
		{"donald", "regular", 100},
		{"scrooge", "vip", 1_000_000},
	}

	// Give VIP accounts 1000 bonus
	/* Value semantics "for" loop
	for _, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
		}
	}
	*/

	// Pointer semantics "for" loop
	for i := range bank {
		if bank[i].Type == "vip" {
			bank[i].Balance += 1_000
			// The compiler will translate the above to this
			//(&bank[i]).Balance += 1_000
		}

	}

	/* 3rd option (transaction) (only option in maps)
	for i, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
		}
		bank[i] = a
	}
	*/

	for _, a := range bank {
		fmt.Println(a)
	}
}
