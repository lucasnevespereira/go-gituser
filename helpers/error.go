package helpers

import (
	"errors"
	"fmt"
	"os"
)

var (
	errInvalidArguments = errors.New("Invalid Arguments 🙁")
	errExecutingMode    = errors.New("Something went wrong executing this mode 😭")
	errReadingInput     = errors.New("Couldn't understand your input 🤯")
)

// PrintErrorInvalidArguments is used to handle error logs for invalid arguments
func PrintErrorInvalidArguments() {
	fmt.Fprintf(os.Stderr, "error: %v \n", errInvalidArguments)
}

// PrintErrorExecutingMode is used to handle error logs after execution
func PrintErrorExecutingMode() {
	fmt.Fprintf(os.Stderr, "error: %v \n", errExecutingMode)
}

// PrintErrorReadingInput is used to handle error logs reading input
func PrintErrorReadingInput() {
	fmt.Fprintf(os.Stderr, "error: %v \n", errReadingInput)
}
