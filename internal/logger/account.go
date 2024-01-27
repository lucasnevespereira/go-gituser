package logger

import (
	"fmt"
	"go-gituser/internal/models"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

func ReadFileData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, errors.Wrap(err, "ReadFileData")
	}
	return data, nil
}

func ReadAccountsData(account *models.Accounts) {
	fmt.Println("Hello, this is your accounts data")
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

func ReadCurrentAccountData(name, email, mode string) {
	fmt.Println("You are on the " + color.CyanString(mode) + " acccount")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", name)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", email)
}

func ReadUnsavedGitAccount(name, email string) {
	fmt.Println("You are using the following account")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", name)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", email)

	fmt.Println("This account is " + color.YellowString("unsaved") + ". Run <gituser setup> to save it to a " + color.CyanString("mode"))
}
