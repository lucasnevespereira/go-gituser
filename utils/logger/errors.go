package logger

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	errInvalidArguments = errors.New("invalid arguments 🙁")
	errExecutingMode    = errors.New("something went wrong executing this mode 😭")
	errReadingInput     = errors.New("couldn't understand your input 🤯")
)

func PrintErrorInvalidArguments() {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", errInvalidArguments)
	fmt.Println("For further information see 'gituser --help'")
}

func PrintErrorExecutingMode() {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", errExecutingMode)
}

func PrintErrorReadingInput() {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", errReadingInput)
}

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", err)
}

func PrintErrorWithMessage(err error, message string) {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+" %s - %v \n", message, err)
}
