package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintManual prints manual
func PrintManual() {
	fmt.Println("Hi there 👋🏼")
	fmt.Println("")
	color.Cyan("Description:")
	fmt.Println("This programs automates the git config command.")
	fmt.Println("There is 3 modes for this program")
	fmt.Println(" - [💻] <work> for a professional account \n - [📚] <school> for a school account \n - [🏠] <personal> for a personal account")
	fmt.Println("")
	color.Cyan("Usage:")
	fmt.Println("To use the program you just need to call the executable")
	fmt.Println("")
	fmt.Println(" gituser <mode>")
	fmt.Println("")
	color.Cyan("Flags:")
	fmt.Println("")
	fmt.Println(" gituser --help (Help Information)")
	fmt.Println(" gituser --manual (Manual Information)")
	fmt.Println(" gituser --info (Accounts information)")
	fmt.Println("")
}
