package main

import "fmt"

type Player struct {
	Name string
	Mana int
}

func main() {
	players := map[string]Player{ // ID -> Player
		"7a": {"Link", 72},
		"3b": {"Zelda", 82},
	}

	for _, p := range players {
		p.Mana += 10 // change a copy, won't reflect in players
	}

	/*
		for k := range players {
			players[k].Mana += 10 // won't compile
		}
	*/

	// Option 1: Change players to map[string]*Player
	// Option 2: Read,Modify,Write (transaction)

	for id, p := range players {
		p.Mana += 10
		players[id] = p
	}

	for _, p := range players {
		fmt.Println(p)
	}
}
