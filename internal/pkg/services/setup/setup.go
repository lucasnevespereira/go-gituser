package setup

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/models"
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
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		switch choice {
		case workSelectLabel:
			s.getUserAccount(models.WorkMode)
		case schoolSelectLabel:
			s.getUserAccount(models.SchoolMode)
		case personalSelectLabel:
			s.getUserAccount(models.PersonalMode)
		case cancelSelectLabel:
			os.Exit(1)
		}

		fmt.Println("Would you like to setup your accounts again? (y/N)")
		shouldConfigureAgain := readInput()
		shouldConfigureAgain = strings.ToUpper(strings.TrimSpace(shouldConfigureAgain))

		if shouldConfigureAgain != strings.ToLower(yes) || shouldConfigureAgain != strings.ToUpper(yes) {
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
		inputWorkUsername = readInput()

		fmt.Println("What is your work email?")
		inputWorkEmail = readInput()

	case models.SchoolMode:
		fmt.Println("What is your school username?")
		inputSchoolUsername = readInput()

		fmt.Println("What is your school email?")
		inputSchoolEmail = readInput()

	case models.PersonalMode:
		fmt.Println("What is your personal username?")
		inputPersonalUsername = readInput()

		fmt.Println("What is your personal email?")
		inputPersonalEmail = readInput()
	}
}

func (s *Service) writeAccountData(account models.Accounts) {

	err := s.saveNonEmptyFields(account)
	if err != nil {
		logger.PrintErrorWithMessage(err, "failed to save account data")
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
		logger.PrintErrorReadingInput()
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
