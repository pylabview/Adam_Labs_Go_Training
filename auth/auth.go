package main

import (
	"fmt"
	"time"
)

func main() {
	// An interface is nil only if both itab and data are nil
	if err := login("bugs", "duckseason"); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("OK")
}

func login(login, passwd string) error {
	// var err *AuthError // BUG: won't be nil error
	var err error

	return err // FIXME
}

type AuthError struct {
	Login  string
	Time   time.Time
	Reason string
}

// Error implement the built-in error interface
func (ue *AuthError) Error() string {
	return fmt.Sprintf("[%s] %q: %s", ue.Time.Format(time.RFC3339), ue.Login, ue.Reason)
	// return "err" // for the benchmark
}
