package helpers

import (
	"errors"
	"fmt"
	"os"
)

var (
	errInvalidArguments = errors.New("Invalid Arguments")
)

// PrintError is used to handle error logs
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "error %v \n", err)
}
