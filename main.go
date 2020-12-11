package main

import (
	"errors"
	"fmt"
	"go-gituser/helpers"
	"os"
	"strings"
)

var (
	errInvalidArguments = errors.New("Invalid Arguments 🙁")
)

func main() {
	if len(os.Args) != 2 {
		helpers.PrintError(errInvalidArguments)
	}

	argValue := strings.ToUpper(os.Args[1])

	switch argValue {
	case "WORK":
		fmt.Println("Work mode")
	case "SCHOOL":
		fmt.Println("School mode")
	case "PERSONAL":
		fmt.Println("Personal mode")
	default:
		fmt.Println("Personal Mode")
	}

}
