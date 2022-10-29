package app

import (
	"encoding/json"
	"fmt"
	"go-gituser/internal/models"
	"go-gituser/state"
	"go-gituser/utils"
	"go-gituser/utils/logger"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

var (
	inputPersonalUsername string
	inputPersonalEmail    string
	inputWorkUsername     string
	inputWorkEmail        string
	inputSchoolUsername   string
	inputSchoolEmail      string
	shouldConfigureAgain  string
)

func SetupAccounts() {

	for {
		prompt := promptui.Select{
			Label: "Please choose an account to configure",
			Items: []string{
				utils.WorkSelectLabel,
				utils.SchoolSelectLabel,
				utils.PersonalSelectLabel,
				utils.CancelSelectLabel,
			},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case utils.WorkSelectLabel:
			getUserAccount(utils.WorkMode)
			logger.PrintRemeberToActiveMode(utils.WorkMode)
		case utils.SchoolSelectLabel:
			getUserAccount(utils.SchoolMode)
			logger.PrintRemeberToActiveMode(utils.SchoolMode)
		case utils.PersonalSelectLabel:
			getUserAccount(utils.PersonalMode)
			logger.PrintRemeberToActiveMode(utils.PersonalMode)
		case utils.CancelSelectLabel:
			os.Exit(1)
		}

		fmt.Println("Would you like to configure another account ? (y/n)")
		_, err = fmt.Scanln(&shouldConfigureAgain)
		if err != nil {
			logger.PrintErrorReadingInput()
		}

		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != utils.Yes {
			fmt.Println("Okay. Bye there!")
			break
		}
	}

	checkForEmptyAccountData()
	writeAccountData()
}

func getUserAccount(mode string) {
	switch mode {
	case utils.WorkMode:
		fmt.Println("What is your work username ?")
		_, errUsername := fmt.Scanln(&inputWorkUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your work email ?")
		_, errEmail := fmt.Scanln(&inputWorkEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

	case utils.SchoolMode:
		fmt.Println("What is your school username ?")
		_, errUsername := fmt.Scanln(&inputSchoolUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your school email ?")
		_, errEmail := fmt.Scanln(&inputSchoolEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}
	case utils.PersonalMode:
		fmt.Println("What is your personal username ?")
		_, errUsername := fmt.Scanln(&inputPersonalUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println("What is your personal email ?")
		_, errEmail := fmt.Scanln(&inputPersonalEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

	case utils.CancelSelectLabel:
		os.Exit(1)
	}

}

func writeAccountData() {
	accountsFile, err := utils.GetAccountsDataFile()
	if err != nil {
		logger.PrintError(err)
	}

	accounts := models.Accounts{
		PersonalUsername: inputPersonalUsername,
		PersonalEmail:    inputPersonalEmail,
		WorkUsername:     inputWorkUsername,
		WorkEmail:        inputWorkEmail,
		SchoolUsername:   inputSchoolUsername,
		SchoolEmail:      inputSchoolEmail,
	}

	file, err := json.MarshalIndent(accounts, "", " ")
	if err != nil {
		logger.PrintErrorWithMessage(err, "json.MarshalIndent")
	}

	os.WriteFile(accountsFile, file, 666)

	state.SavedAccounts = &accounts
}

// checkForEmptyAccountData checks if there is no overrides with empty accounts.
func checkForEmptyAccountData() {
	Sync()

	if inputPersonalEmail == "" || inputPersonalUsername == "" {
		inputPersonalEmail = state.SavedAccounts.PersonalEmail
		inputPersonalUsername = state.SavedAccounts.PersonalUsername
	}

	if inputWorkEmail == "" || inputWorkUsername == "" {
		inputWorkEmail = state.SavedAccounts.WorkEmail
		inputWorkUsername = state.SavedAccounts.WorkUsername
	}

	if inputSchoolEmail == "" || inputSchoolUsername == "" {
		inputSchoolEmail = state.SavedAccounts.SchoolEmail
		inputSchoolUsername = state.SavedAccounts.SchoolUsername
	}

}
