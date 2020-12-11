package helpers

import "fmt"

// PrintHelp prints help
func PrintHelp() {
	fmt.Println("Hi there 👋🏼")
	fmt.Println("This is the help manual")
	fmt.Println("")
	fmt.Println("Description:")
	fmt.Println("This programs automates the git config command.")
	fmt.Println("There is 3 modes for this program")
	fmt.Println(" - <work> for a professional account \n - <school> for a school account \n - <personal> for a personal account")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("To use the program you just need to call the executable")
	fmt.Println("")
	fmt.Println(" ./gituser <mode>")
	fmt.Println("")
}
