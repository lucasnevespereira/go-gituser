package helpers

import (
	"fmt"
	"go-gituser/models"
	"io/ioutil"

	"github.com/fatih/color"
)

// GetDataFromJSON is a func that returns the data from a JSON File
func GetDataFromJSON(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}

// ReadAccountsData prints accounts infos to the user
func ReadAccountsData(account models.Account) {
	fmt.Println("Hello, this is your accounts data 🧑🏻‍💻")
	fmt.Println("")
	if account.PersonalUsername == "" {
		fmt.Println("🏠 | You have no personal account defined")
	} else {
		fmt.Println("🏠 | Personal Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.PersonalUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.PersonalEmail)
	}
	fmt.Println("")
	if account.SchoolUsername == "" {
		fmt.Println("📚 | You have no school account defined")
	} else {
		fmt.Println("📚 | School Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.SchoolUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.SchoolEmail)
	}
	fmt.Println("")
	if account.WorkUsername == "" {
		fmt.Println("💻 | You have no work account defined")
	} else {
		fmt.Println("💻 | Work Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.WorkUsername)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.WorkEmail)
	}
	fmt.Println("")

}

// ReadCurrentAccountData prints current account data
func ReadCurrentAccountData(name string, email string, mode string) {
	fmt.Println("You are on the " + color.CyanString(mode) + " acccount")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", name)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", email)
}
