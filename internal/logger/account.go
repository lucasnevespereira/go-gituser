package logger

import (
	"fmt"
	"go-gituser/internal/models"

	"github.com/fatih/color"
)

func ReadAccountsData(accounts *models.Accounts) {
	fmt.Println("Hello, this is your accounts data")
	fmt.Println("")
	if accounts.Personal.Username == "" {
		fmt.Println("🏠 | You have no personal account defined")
	} else {
		fmt.Println("🏠 | Personal Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", accounts.Personal.Username)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", accounts.Personal.Email)
		if accounts.Personal.SigningKeyID != "" {
			fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", accounts.Personal.SigningKeyID)
		}
		if accounts.Personal.SSHKeyPath != "" {
			fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", accounts.Personal.SSHKeyPath)
		}
	}
	fmt.Println("")
	if accounts.School.Username == "" {
		fmt.Println("📚 | You have no school account defined")
	} else {
		fmt.Println("📚 | School Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", accounts.School.Username)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", accounts.School.Email)
		if accounts.School.SigningKeyID != "" {
			fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", accounts.School.SigningKeyID)
		}
		if accounts.School.SSHKeyPath != "" {
			fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", accounts.School.SSHKeyPath)
		}

	}
	fmt.Println("")
	if accounts.Work.Username == "" {
		fmt.Println("💻 | You have no work account defined")
	} else {
		fmt.Println("💻 | Work Git Account :")
		fmt.Printf(color.BlueString("=>")+" Username: %v\n", accounts.Work.Username)
		fmt.Printf(color.BlueString("=>")+" Email: %v\n", accounts.Work.Email)
		if accounts.Work.SigningKeyID != "" {
			fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", accounts.Work.SigningKeyID)
		}
		if accounts.Work.SSHKeyPath != "" {
			fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", accounts.Work.SSHKeyPath)
		}

	}
	fmt.Println("")

}

func ReadCurrentAccountData(account *models.Account, mode string) {
	fmt.Println("You are on the " + color.CyanString(mode) + " acccount")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.Username)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.Email)
	if account.SigningKeyID != "" {
		fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", account.SigningKeyID)
	}
	if account.SSHKeyPath != "" {
		fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", account.SSHKeyPath)
	}
}

func ReadUnsavedGitAccount(account *models.Account) {
	fmt.Println("You are using the following account")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.Username)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.Email)
	if account.SigningKeyID != "" {
		fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", account.SigningKeyID)
	}
	if account.SSHKeyPath != "" {
		fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", account.SSHKeyPath)
	}

	fmt.Println("This account is " + color.YellowString("unsaved") + ". Run <gituser setup> to save it to a " + color.CyanString("mode"))
}
