package goBasic

import "fmt"

/*
 * The error type is a built-in interface
 * type error interface {
 *	 Error() string
 * }
 */

type MyError struct {
	message string
}

// The MyError matches with the built-in error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %v", e.message)
}

func run() error {
	return &MyError{
		message: "Internal server error",
	}
}

func PlayWithError() {
	errMsg := run()
	fmt.Println(errMsg)
}