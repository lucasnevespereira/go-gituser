package main

import (
	"errors"
	"go-gituser/helpers"
	"os"
)

var (
	errInvalidArguments = errors.New("Invalid Arguments 🙁")
)

func main() {
	if len(os.Args) != 2 {
		helpers.PrintError(errInvalidArguments)
	}
}
