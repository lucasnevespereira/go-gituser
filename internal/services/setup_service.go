package services

import (
	"fmt"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type ISetupService interface {
	SetupAccounts() error
}

type SetupService struct {
	accountService IAccountService
}

func NewSetupService(accountService IAccountService) ISetupService {
	return &SetupService{
		accountService: accountService,
	}
}

var (
	inputPersonalUsername string
	inputPersonalEmail    string
	inputWorkUsername     string
	inputWorkEmail        string
	inputSchoolUsername   string
	inputSchoolEmail      string
	shouldConfigureAgain  string
)

const (
	workSelectLabel     = "💻 Work Account"
	schoolSelectLabel   = "📚 School Account"
	personalSelectLabel = "🏠 Personal Account"
	cancelSelectLabel   = "Cancel"
	yes                 = "Y"
)

func (s *SetupService) SetupAccounts() error {
	for {
		prompt := promptui.Select{
			Label: "Please choose an account to configure",
			Items: []string{
				workSelectLabel,
				schoolSelectLabel,
				personalSelectLabel,
				cancelSelectLabel,
			},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case workSelectLabel:
			selectUserAccount(models.WorkMode)
			logger.PrintRemeberToActiveMode(models.WorkMode)
		case schoolSelectLabel:
			selectUserAccount(models.SchoolMode)
			logger.PrintRemeberToActiveMode(models.SchoolMode)
		case personalSelectLabel:
			selectUserAccount(models.PersonalMode)
			logger.PrintRemeberToActiveMode(models.PersonalMode)
		case cancelSelectLabel:
			os.Exit(1)
		}

		fmt.Println("Would you like to configure another account ? (y/n)")
		_, err = fmt.Scanln(&shouldConfigureAgain)
		if err != nil {
			logger.PrintErrorReadingInput()
		}

		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != yes {
			fmt.Println("Okay. Bye there!")
			break
		}
	}
	savedAccounts, err := s.accountService.ReadSavedAccounts()
	if err != nil {
		return models.ErrSetupAccounts
	}

	checkForEmptyAccountData(savedAccounts)
	if err = s.accountService.SaveAccounts(&models.Accounts{
		PersonalUsername: inputPersonalUsername,
		PersonalEmail:    inputPersonalEmail,
		WorkUsername:     inputWorkUsername,
		WorkEmail:        inputWorkEmail,
		SchoolUsername:   inputSchoolUsername,
		SchoolEmail:      inputSchoolEmail,
	}); err != nil {
		return models.ErrSetupAccounts
	}

	return nil
}

func selectUserAccount(mode string) {
	switch mode {
	case models.WorkMode:
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

	case models.SchoolMode:
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
	case models.PersonalMode:
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

	case cancelSelectLabel:
		os.Exit(1)
	}

}

// checkForEmptyAccountData checks if there is no overrides with empty accounts.
func checkForEmptyAccountData(savedAccounts *models.Accounts) {

	if inputPersonalEmail == "" || inputPersonalUsername == "" {
		inputPersonalEmail = savedAccounts.PersonalEmail
		inputPersonalUsername = savedAccounts.PersonalUsername
	}

	if inputWorkEmail == "" || inputWorkUsername == "" {
		inputWorkEmail = savedAccounts.WorkEmail
		inputWorkUsername = savedAccounts.WorkUsername
	}

	if inputSchoolEmail == "" || inputSchoolUsername == "" {
		inputSchoolEmail = savedAccounts.SchoolEmail
		inputSchoolUsername = savedAccounts.SchoolUsername
	}
}
