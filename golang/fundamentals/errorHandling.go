package main

import (
	"errors"
	"fmt"
)

// Create a predefined/sentinel error.
// We can compare other errors against this specific error.
var ErrNotFound = errors.New("user not found")

func getUser() error {

	// %w wraps ErrNotFound inside a new error.
	//
	// The final error message becomes:
	// "database lookup failed: user not found"
	//
	// But Go also remembers that ErrNotFound
	// is the underlying/original error.
	return fmt.Errorf("database lookup failed: %w", ErrNotFound)
}

func main() {

	// Call getUser().
	//
	// err contains the wrapped error:
	//
	// database lookup failed: user not found
	err := getUser()

	// Print the error message.
	fmt.Println(err)

	// errors.Is() checks whether err is:
	// 1. exactly ErrNotFound, OR
	// 2. an error that wraps ErrNotFound somewhere in its chain.

	// Because getUser() used %w, ErrNotFound is
	// part of the error chain.
	if errors.Is(err, ErrNotFound) {
		fmt.Println("The actual problem was: user not found")
	}
}
