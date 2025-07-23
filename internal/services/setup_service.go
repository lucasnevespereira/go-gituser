package services

import (
	"bufio"
	"fmt"
	"go-gituser/internal/connectors/ssh"
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
	inputPersonalUsername     string
	inputPersonalEmail        string
	inputPersonalSigningKeyID string
	inputPersonalSSHKeyPath   string
	inputWorkUsername         string
	inputWorkEmail            string
	inputWorkSigningKeyID     string
	inputWorkSSHKeyPath       string
	inputSchoolUsername       string
	inputSchoolEmail          string
	inputSchoolSigningKeyID   string
	inputSchoolSSHKeyPath     string
	shouldConfigureAgain      string
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
		Personal: models.Account{
			Username:     inputPersonalUsername,
			Email:        inputPersonalEmail,
			SigningKeyID: inputPersonalSigningKeyID,
		},
		Work: models.Account{
			Username:     inputWorkUsername,
			Email:        inputWorkEmail,
			SigningKeyID: inputWorkSigningKeyID,
		},
		School: models.Account{
			Username:     inputSchoolUsername,
			Email:        inputSchoolEmail,
			SigningKeyID: inputSchoolSigningKeyID,
		},
	}); err != nil {
		return models.ErrSetupAccounts
	}

	return nil
}

func askForSSHKey(sshConnector ssh.ISSHConnector) string {
	var useSSH string
	fmt.Println("Would you like to configure an SSH key for this account? (y/n)")
	_, err := fmt.Scanln(&useSSH)
	if err != nil {
		logger.PrintErrorReadingInput()
		os.Exit(1)
	}

	if strings.ToUpper(strings.TrimSpace(useSSH)) != yes {
		return ""
	}

	var sshKeyPath string
	defaultPath := sshConnector.GetDefaultKeyPath()
	fmt.Printf("Enter the path to your SSH private key (default: %s): ", defaultPath)

	// Read the full line to handle paths with spaces
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		sshKeyPath = strings.TrimSpace(scanner.Text())
	}

	if sshKeyPath == "" {
		sshKeyPath = defaultPath
	}

	// Validate the SSH key
	if err := sshConnector.ValidateKeyPath(sshKeyPath); err != nil {
		fmt.Printf("⚠️  Warning: %v\n", err)
		fmt.Println("You can update this later by running setup again.")
	}

	return sshKeyPath
}

func askForGPGKey() bool {
	var useGPG string
	fmt.Println("Would you like to use GPG signing for this account? (y/n)")
	_, err := fmt.Scanln(&useGPG)
	if err != nil {
		logger.PrintErrorReadingInput()
		os.Exit(1)
	}
	return strings.ToUpper(strings.TrimSpace(useGPG)) == yes
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

		if askForGPGKey() {
			fmt.Println("What is your work gpg signing key id ?")
			_, errSigningKeyID := fmt.Scanln(&inputWorkSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
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

		if askForGPGKey() {
			fmt.Println("What is your school gpg signing key id ?")
			_, errSigningKeyID := fmt.Scanln(&inputSchoolSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
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

		if askForGPGKey() {
			fmt.Println("What is your personal gpg signing key id ?")
			_, errSigningKeyID := fmt.Scanln(&inputPersonalSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
		}

	case cancelSelectLabel:
		os.Exit(1)
	}

}

func checkForEmptyAccountData(savedAccounts *models.Accounts) {
	if inputPersonalEmail == "" || inputPersonalUsername == "" {
		inputPersonalEmail = savedAccounts.Personal.Email
		inputPersonalUsername = savedAccounts.Personal.Username
		if inputPersonalSigningKeyID == "" {
			inputPersonalSigningKeyID = savedAccounts.Personal.SigningKeyID
		}
	}

	if inputWorkEmail == "" || inputWorkUsername == "" {
		inputWorkEmail = savedAccounts.Work.Email
		inputWorkUsername = savedAccounts.Work.Username
		if inputWorkSigningKeyID == "" {
			inputWorkSigningKeyID = savedAccounts.Work.SigningKeyID
		}
	}

	if inputSchoolEmail == "" || inputSchoolUsername == "" {
		inputSchoolEmail = savedAccounts.School.Email
		inputSchoolUsername = savedAccounts.School.Username
		if inputSchoolSigningKeyID == "" {
			inputSchoolSigningKeyID = savedAccounts.School.SigningKeyID
		}
	}
}
