package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintHelp prints help
func PrintHelp() {
	fmt.Println("Hi there 👋🏼")
	fmt.Println("This is the help manual")
	fmt.Println("")
	color.Cyan("Description:")
	fmt.Println("This program helps switch between different git user accounts easily")
	fmt.Println("There is 3 modes for this program")
	fmt.Println(" - 🏠 <personal> for a personal account \n - 📚 <school> for a school account \n - 💻 <work> for a professional account")
	fmt.Println("")
	color.Cyan("Usage:")
	fmt.Println("To use the program you just need to call the executable")
	fmt.Println(" gituser <mode>")
	fmt.Println("")
	color.Cyan("Flags:")
	fmt.Println(" gituser --help (Help manual)")
	fmt.Println(" gituser --info (Accounts information)")
	fmt.Println("")
}
