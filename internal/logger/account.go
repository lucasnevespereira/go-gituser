package logger

import (
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"strings"

	"github.com/fatih/color"
)

func ReadAccountsData(accounts *models.Accounts) {
	shown := false
	accounts.ForEachConfigured(func(mode string, account models.Account) bool {
		if !shown {
			fmt.Println("Hello, this is your accounts data")
		}
		fmt.Println()
		switch mode {
		case models.PersonalMode:
			fmt.Println("🏠 | Personal Git Account :")
		case models.SchoolMode:
			fmt.Println("📚 | School Git Account :")
		case models.WorkMode:
			fmt.Println("💻 | Work Git Account :")
		default:
			fmt.Printf("🔖 | %s Git Account :\n", mode)
		}
		printSavedAccountFields(account)
		shown = true
		return true
	})
	if !shown {
		fmt.Println("No accounts configured. Run gituser setup to add one.")
	}
}

func printSavedAccountFields(account models.Account) {
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.Username)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.Email)
	if account.SigningKeyID != "" {
		fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", account.SigningKeyID)
	}
	if account.SSHKeyPath != "" {
		fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", account.SSHKeyPath)
	}
}

func ReadCurrentAccountData(account *models.Account, mode string) {
	fmt.Println("You are on the " + color.CyanString(mode) + " account")
	fmt.Printf(color.BlueString("=>")+" Username: %v\n", account.Username)
	fmt.Printf(color.BlueString("=>")+" Email: %v\n", account.Email)
	if account.SigningKeyID != "" {
		fmt.Printf(color.BlueString("=>")+" Signing Key ID: %v\n", account.SigningKeyID)
	}
	if account.SSHKeyPath != "" {
		fmt.Printf(color.BlueString("=>")+" SSH Key: %v\n", account.SSHKeyPath)
	}
}

func ReadMatchingAccountsData(account *models.Account, modes []string) {
	fmt.Println("This Git account matches multiple saved modes: " + color.CyanString(strings.Join(modes, ", ")))
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
