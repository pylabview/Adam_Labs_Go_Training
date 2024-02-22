package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`1`)

	// int8, int16, int32, int64, int, uint8, ..., uint, float32, float64
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(v)

	// v := json.Unmarshal(data) // float64

	// zero vs missing: Only at API & storage layers
	var event struct { // anonymous struct
		User  string
		Likes *int
	}

	data = []byte(`{"user": "joe", "likes": 0}`)
	if err := json.Unmarshal(data, &event); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%#v\n", event)
	if event.Likes == nil {
		fmt.Println("missing likes")
	} else {
		fmt.Printf("%d likes\n", *event.Likes)
	}

}

// API -> Business -> Storage
