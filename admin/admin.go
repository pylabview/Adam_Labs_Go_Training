package main

import (
	"fmt"
)

func main() {
	u := User{"joe", "joe@e-corp.com"}
	u.Notify("Submit your time sheet")

	a := Admin{
		User:  User{"jess", "jess@e-corp.com"},
		Level: "supervisor",
	}
	a.Notify("upgrade firewall")
	// a.User.Notify()
	// a.Name
	// a.User.Name

	// notify(u, "update OS") // won't compile (method)
	notify(&u, "update OS")
	notify(&a, "check logs")
}

func notify(n Notifier, msg string) {
	n.Notify(msg)
}

type Notifier interface {
	Notify(string) error
}

type Admin struct {
	User  // Admin embeds User
	Level string
}

// *: a pointer type (or converting from a pointer to a value, rare)
// &: take address of value, value -> pointer

// Traditional OO: this, self can be this class or any class inheriting from it
// In Go: The receiver (u) has a specific type
func (u *User) Notify(msg string) error {
	fmt.Printf("%s <%s>: %s\n", u.Name, u.Email, msg)
	return nil
}

type User struct {
	Name  string
	Email string
}
