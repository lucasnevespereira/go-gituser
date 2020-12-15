package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	errInvalidArguments = errors.New("Invalid Arguments 🙁")
	errExecutingMode    = errors.New("Something went wrong executing this mode 😭")
	errReadingInput     = errors.New("Couldn't understand your input 🤯")
)

// PrintErrorInvalidArguments is used to handle error logs for invalid arguments
func PrintErrorInvalidArguments() {
	fmt.Fprintf(os.Stderr, color.RedString("Error: ")+"%v \n", errInvalidArguments)
	fmt.Println("For further information see 'gituser --help'")
}

// PrintErrorExecutingMode is used to handle error logs after execution
func PrintErrorExecutingMode() {
	fmt.Fprintf(os.Stderr, color.RedString("Error: ")+"%v \n", errExecutingMode)
}

// PrintErrorReadingInput is used to handle error logs reading input
func PrintErrorReadingInput() {
	fmt.Fprintf(os.Stderr, color.RedString("Error: ")+"%v \n", errReadingInput)
}

// PrintErrorReadingAccount is used to handle error logs reading accounts
func PrintErrorReadingAccount(mode string) {
	fmt.Println("")
	fmt.Printf("You have no %v account! \n", mode)
	fmt.Println("")
	color.Cyan("Additional info:")
	fmt.Printf("To add a %v account you need to add it to data/config.json \n", mode)
	fmt.Println("Then recompile your program:")
	fmt.Println("Step 1 : Go to the source directory for this project")
	fmt.Println("Step 2 : Once there please run go build -o gituser")
	fmt.Println("")
}
