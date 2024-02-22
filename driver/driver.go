// Example how traditional OO doesn't work in Go
package main

import (
	"errors"
	"fmt"
)

func main() {
	var d PGDriver
	fmt.Println(d.Query("SELECT user, balance FROM users"))
}

func (d *PGDriver) query(query string) ([]Row, error) {
	return []Row{}, nil
}

type PGDriver struct {
	Driver
}

func (d *Driver) query(query string) ([]Row, error) {
	return nil, ErrNotImplemented
}

func (d *Driver) Query(query string) ([]Row, error) {
	if query == "" {
		return nil, fmt.Errorf("empty query")
	}

	return d.query(query)
}

type Driver struct{}

type Row struct{}

var ErrNotImplemented = errors.New("not implemented")
