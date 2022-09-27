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

// PrintError is used to handle error logs
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, color.RedString("Error: ")+"%v \n", err)
}

// PrintWarningReadingAccount is used to handle warning logs reading accounts
func PrintWarningReadingAccount(mode string) {
	fmt.Println("")
	fmt.Fprintf(os.Stderr, color.YellowString("Warning: ")+"%v \n", "You have no "+mode+" account 🧐")
	fmt.Println("")
	color.Cyan("Tips:")
	fmt.Printf("To add a %v account try to run gituser config \n", mode)
	fmt.Println("")
}

// PrintRemeberToActiveMode is used to remember user to active the mode after config
func PrintRemeberToActiveMode(mode string) {
	info := color.New(color.Bold).PrintfFunc()
	info(color.BlueString("Remember to run <gituser %v> to activate this mode \n"), mode)
}
