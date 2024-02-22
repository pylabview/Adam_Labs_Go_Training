/*
ETL: Extract, transform, load
Always start with concreted data types, then move to interfaces
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

/* 3 layers code
primitive: Victoria.Pull, Henry.Store
lower: pull, store
higher: Copy
*/

type Data struct {
	Line string
}

type Victoria struct {
	Host string

	n int
}

// func (v *Victoria) Pull() (Data, error) {
func (v *Victoria) Pull(d *Data) error {
	v.n++

	if v.n > 10 {
		return io.EOF
	}

	d.Line = fmt.Sprintf("data #%d", v.n)
	log.Printf("Victoria: %s", d.Line)
	return nil
}

type Ann struct {
	Host string

	n int
}

func (v *Ann) Pull(d *Data) error {
	v.n++

	if v.n > 10 {
		return io.EOF
	}

	d.Line = fmt.Sprintf("data #%d", v.n)
	log.Printf("Ann: %s", d.Line)
	return nil
}

type Puller interface {
	Pull(*Data) error
}

type Henry struct {
	Host string
}

func (h *Henry) Store(d *Data) error {
	log.Printf("Henry: %s", d.Line)
	return nil
}

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		err := p.Pull(&data[i])
		if errors.Is(err, io.EOF) {
			return i, nil
		}

		if err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func store(h *Henry, data []Data) (int, error) {
	for i := range data {
		if err := h.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func Copy(m *Henry, p Puller, batchSize int) error {
	data := make([]Data, batchSize)

	for {
		i, err := pull(p, data)
		if i == 0 {
			break
		}

		if err != nil {
			return err
		}

		if _, err := store(m, data[:i]); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	/*
		v := Victoria{
			Host: "localhost:8000",
		}
	*/
	p := Ann{
		Host: "localhost:8000",
	}

	m := Henry{
		Host: "localhost:9000",
	}

	if err := Copy(&m, &p, 3); err != io.EOF {
		fmt.Println(err)
	}
}

/* Interfaces rules of thumb:
- Accept interface, return types
- Small interfaces (< 5)
- Think about data semantics
- Use cases
	- Polymorphism
	- Change how stdlib looks at your types
		fmt.Stringer
		json.Marshaler, json.Unmarshaler
	- Mocking (try avoid, docker ...)
*/
