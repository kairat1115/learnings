package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Query string
	Err   error
}

func (this *QueryError) Error() string {
	return this.Query + ": " + this.Err.Error()
}

// this is new in go 1.13
func (this *QueryError) Unwrap() error {
	return this.Err
}

var ErrNotFound = errors.New("not found")
var ErrPermission = errors.New("permission denied")

func main() {
	var err error = nil
	// Similar to:
	//   if err == ErrNotFound {...}
	if errors.Is(err, ErrNotFound) {
		// something was not found
	}

	// Similar to:
	//   if e, ok := err.(*QueryError); ok { … }
	var e *QueryError
	// Note: *QueryError is the type of the error
	if errors.As(err, &e) {
		// err is a *QueryError, and e is set to the error's value
	}

	if errors.Is(err, ErrPermission) {
		// err, or some error that it wraps, is a permission problem
	}

	name := "Bob"
	if err != nil {
		// Return an error which unwraps to err.
		fmt.Errorf("decompress %v: %w", name, err)
	}
}
