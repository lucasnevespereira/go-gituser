package setup

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"go-gituser/internal/pkg/external/repository"
	logger2 "go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/models"
	"go-gituser/state"
	"os"
	"strings"
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

const (
	workSelectLabel     = "💻 Work Account"
	schoolSelectLabel   = "📚 School Account"
	personalSelectLabel = "🏠 Personal Account"
	cancelSelectLabel   = "Cancel"
	yes                 = "Y"
	no                  = "N"
)

type Service struct {
	repository repository.AccountRepository
}

func NewSetupService(repository repository.AccountRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) SetupAccounts() {
	for {
		choice, err := s.getAccountChoice()
		if err != nil {
			logger2.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case workSelectLabel:
			s.getUserAccount(models.WorkMode)
			logger2.PrintRemeberToActiveMode(models.WorkMode)
		case schoolSelectLabel:
			s.getUserAccount(models.SchoolMode)
			logger2.PrintRemeberToActiveMode(models.SchoolMode)
		case personalSelectLabel:
			s.getUserAccount(models.PersonalMode)
			logger2.PrintRemeberToActiveMode(models.PersonalMode)
		case cancelSelectLabel:
			os.Exit(1)
		}

		fmt.Println("Would you like to setup another account? (y/n)")
		shouldConfigureAgain := readInput()
		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != yes {
			fmt.Println("Okay. Bye there!")
			break
		}
	}

	accounts := models.Accounts{
		PersonalUsername: inputPersonalUsername,
		PersonalEmail:    inputPersonalEmail,
		WorkUsername:     inputWorkUsername,
		WorkEmail:        inputWorkEmail,
		SchoolUsername:   inputSchoolUsername,
		SchoolEmail:      inputSchoolEmail,
	}

	s.writeAccountData(accounts)
}

func (s *Service) getAccountChoice() (string, error) {
	prompt := promptui.Select{
		Label: "Please choose an account to setup",
		Items: []string{
			workSelectLabel,
			schoolSelectLabel,
			personalSelectLabel,
			cancelSelectLabel,
		},
	}

	_, choice, err := prompt.Run()
	return choice, err
}

func (s *Service) getUserAccount(mode string) {
	switch mode {
	case models.WorkMode:
		fmt.Println("What is your work username?")
		inputWorkUsername := readInput()
		if inputWorkUsername != "" {
			state.SavedAccounts.WorkUsername = inputWorkUsername
		}

		fmt.Println("What is your work email?")
		inputWorkEmail := readInput()
		if inputWorkEmail != "" {
			state.SavedAccounts.WorkEmail = inputWorkEmail
		}

	case models.SchoolMode:
		fmt.Println("What is your school username?")
		inputSchoolUsername := readInput()
		if inputSchoolUsername != "" {
			state.SavedAccounts.SchoolUsername = inputSchoolUsername
		}

		fmt.Println("What is your school email?")
		inputSchoolEmail := readInput()
		if inputSchoolEmail != "" {
			state.SavedAccounts.SchoolEmail = inputSchoolEmail
		}

	case models.PersonalMode:
		fmt.Println("What is your personal username?")
		inputPersonalUsername := readInput()
		if inputPersonalUsername != "" {
			state.SavedAccounts.PersonalUsername = inputPersonalUsername
		}

		fmt.Println("What is your personal email?")
		inputPersonalEmail := readInput()
		if inputPersonalEmail != "" {
			state.SavedAccounts.PersonalEmail = inputPersonalEmail
		}
	}
}

func (s *Service) writeAccountData(account models.Accounts) {

	err := s.saveNonEmptyFields(account)
	if err != nil {
		logger2.PrintErrorWithMessage(err, "failed to save account data")
	}
}

func (s *Service) saveNonEmptyFields(account models.Accounts) error {
	if account.WorkUsername != "" {
		err := s.repository.SaveAccounts(models.Accounts{WorkUsername: account.WorkUsername})
		if err != nil {
			return err
		}
	}

	if account.WorkEmail != "" {
		err := s.repository.SaveAccounts(models.Accounts{WorkEmail: account.WorkEmail})
		if err != nil {
			return err
		}
	}

	if account.SchoolUsername != "" {
		err := s.repository.SaveAccounts(models.Accounts{SchoolUsername: account.SchoolUsername})
		if err != nil {
			return err
		}
	}

	if account.SchoolEmail != "" {
		err := s.repository.SaveAccounts(models.Accounts{SchoolEmail: account.SchoolEmail})
		if err != nil {
			return err
		}
	}

	if account.PersonalUsername != "" {
		err := s.repository.SaveAccounts(models.Accounts{PersonalUsername: account.PersonalUsername})
		if err != nil {
			return err
		}
	}

	if account.PersonalEmail != "" {
		err := s.repository.SaveAccounts(models.Accounts{PersonalEmail: account.PersonalEmail})
		if err != nil {
			return err
		}
	}

	return nil
}

func readInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		logger2.PrintErrorReadingInput()
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
