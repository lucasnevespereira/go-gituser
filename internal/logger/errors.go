package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func PrintErrorExecutingMode() {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", "something went wrong executing this mode 😭")
}

func PrintErrorReadingInput() {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", "couldn't understand your input 🤯")
}

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+"%v \n", err)
}

func PrintErrorWithMessage(err error, message string) {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+" %s - %v \n", message, err)
}
