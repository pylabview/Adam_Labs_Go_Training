package user

type User struct {
	Name  string
	Email string
}

// New(name, email) (*User, error)
// func New(name, email string) *User {
// func New(name, email) (User, error)
func New(name, email string) *User {
	u := User{
		Name:  name,
		Email: email,
	}

	return &u // The go compiler does escape analysis and will allocate u on the heap
}

/*
// value
func (u User) Update() User {}

// pointer
func (u *User) Update() {}
*/

// go build -gcflags=-m
// escape analysis:
// Does the value outlive the function that owns it? yes -> heap, no -> stack

/* Value vs Pointer semantics
- Rule of thumb: Use value semantics
- Do you want each function to have its own copy: value
- Do you want to share: pointer
*/

/* Go's GC
- mark & sweep
- no compacting
	- nothing  moves
- non generational
	- most objects die young
- concurrent
	- runs alongside your code
	- Mark start (STW < 100us), Marking (concurrent), Mark termination (STW)
- tri color
	- Black, gray, white
*/
