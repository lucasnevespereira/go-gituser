package helpers

import (
	"fmt"
	"go-gituser/models"
	"io/ioutil"
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
		fmt.Println("You have no personal account defined !")
	} else {
		fmt.Println("[🏠] Personal Git Account :")
		fmt.Printf("Username: %v\n", account.PersonalUsername)
		fmt.Printf("Email: %v\n", account.PersonalEmail)
	}
	fmt.Println("")
	if account.SchoolUsername == "" {
		fmt.Println("You have no school account defined !")
	} else {
		fmt.Println("[📚] School Git Account :")
		fmt.Printf("Username: %v\n", account.SchoolUsername)
		fmt.Printf("Email: %v\n", account.SchoolEmail)
	}
	fmt.Println("")
	if account.WorkUsername == "" {
		fmt.Println("You have no work account defined !")
	} else {
		fmt.Println("[💻] Work Git Account :")
		fmt.Printf("Username: %v\n", account.WorkUsername)
		fmt.Printf("Email: %v\n", account.WorkUsername)
	}
	fmt.Println("")

}
