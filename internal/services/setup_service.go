package services

import (
	"bufio"
	"fmt"
	"go-gituser/internal/connectors/ssh"
	"go-gituser/internal/format"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"os"
	"path/filepath"
	"strconv"
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
			s.selectUserAccount(models.WorkMode)
			logger.PrintRemeberToActiveMode(models.WorkMode)
		case schoolSelectLabel:
			s.selectUserAccount(models.SchoolMode)
			logger.PrintRemeberToActiveMode(models.SchoolMode)
		case personalSelectLabel:
			s.selectUserAccount(models.PersonalMode)
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

	savedAccounts, err := s.accountService.GetSavedAccounts()
	if err != nil {
		return models.ErrSetupAccounts
	}

	s.checkForEmptyAccountData(savedAccounts)

	if err = s.accountService.SaveAccounts(&models.Accounts{
		Personal: models.Account{
			Username:     inputPersonalUsername,
			Email:        inputPersonalEmail,
			SigningKeyID: inputPersonalSigningKeyID,
			SSHKeyPath:   inputPersonalSSHKeyPath,
		},
		Work: models.Account{
			Username:     inputWorkUsername,
			Email:        inputWorkEmail,
			SigningKeyID: inputWorkSigningKeyID,
			SSHKeyPath:   inputWorkSSHKeyPath,
		},
		School: models.Account{
			Username:     inputSchoolUsername,
			Email:        inputSchoolEmail,
			SigningKeyID: inputSchoolSigningKeyID,
			SSHKeyPath:   inputSchoolSSHKeyPath,
		},
	}); err != nil {
		return models.ErrSetupAccounts
	}

	return nil
}

func (s *SetupService) selectUserAccount(mode string) {
	// Create SSH connectors for this setup session
	sshConnector := ssh.NewSSHConnector()
	sshDiscovery := NewSSHDiscoveryService(sshConnector)

	switch mode {
	case models.WorkMode:
		fmt.Println("\n=== 💻 Work Account Setup ===")
		fmt.Println("What is your work username?")
		_, errUsername := fmt.Scanln(&inputWorkUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("What is your work email?")
		_, errEmail := fmt.Scanln(&inputWorkEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}
		fmt.Println()
		if s.askForGPGKey(mode) {
			fmt.Println("What is your work GPG signing key ID?")
			_, errSigningKeyID := fmt.Scanln(&inputWorkSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
		}
		// SSH setup
		fmt.Println()
		inputWorkSSHKeyPath = s.setupSSHKeyForAccount("work", inputWorkEmail, sshDiscovery)
		fmt.Println()

	case models.SchoolMode:
		fmt.Println("\n=== 📚 School Account Setup ===")
		fmt.Println("What is your school username?")
		_, errUsername := fmt.Scanln(&inputSchoolUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println()
		fmt.Println("What is your school email?")
		_, errEmail := fmt.Scanln(&inputSchoolEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println()
		if s.askForGPGKey(mode) {
			fmt.Println("What is your school GPG signing key ID?")
			_, errSigningKeyID := fmt.Scanln(&inputSchoolSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
		}

		// SSH setup
		fmt.Println()
		inputSchoolSSHKeyPath = s.setupSSHKeyForAccount("school", inputSchoolEmail, sshDiscovery)
		fmt.Println()

	case models.PersonalMode:
		fmt.Println("\n=== 🏠 Personal Account Setup ===")
		fmt.Println("What is your personal username?")
		_, errUsername := fmt.Scanln(&inputPersonalUsername)
		if errUsername != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("What is your personal email?")
		_, errEmail := fmt.Scanln(&inputPersonalEmail)
		if errEmail != nil {
			logger.PrintErrorReadingInput()
			os.Exit(1)
		}

		fmt.Println()
		if s.askForGPGKey(mode) {
			fmt.Println("What is your personal GPG signing key ID?")
			_, errSigningKeyID := fmt.Scanln(&inputPersonalSigningKeyID)
			if errSigningKeyID != nil {
				logger.PrintErrorReadingInput()
				os.Exit(1)
			}
		}

		// SSH setup
		fmt.Println()
		inputPersonalSSHKeyPath = s.setupSSHKeyForAccount("personal", inputPersonalEmail, sshDiscovery)
		fmt.Println()

	case cancelSelectLabel:
		os.Exit(1)
	}
}

func (s *SetupService) askForGPGKey(mode string) bool {
	fmt.Printf("\n🔑 GPG Key Setup for %s Account\n", format.TitleCase(mode))
	fmt.Println("=====================================")

	var useGPG string
	fmt.Println("Would you like to use GPG signing for this account? (y/n)")
	_, err := fmt.Scanln(&useGPG)
	if err != nil {
		logger.PrintErrorReadingInput()
		os.Exit(1)
	}
	return strings.ToUpper(strings.TrimSpace(useGPG)) == yes
}

func (s *SetupService) setupSSHKeyForAccount(mode, email string, sshDiscovery ISSHDiscoveryService) string {
	fmt.Printf("\n🔑 SSH Key Setup for %s Account\n", format.TitleCase(mode))
	fmt.Println("=====================================")

	// Ask if user wants to setup SSH
	var wantsSSH string
	fmt.Println("Would you like to configure an SSH key for this account? (y/n)")
	fmt.Println("💡 SSH keys allow secure authentication with GitHub/GitLab without passwords")
	_, err := fmt.Scanln(&wantsSSH)
	if err != nil {
		logger.PrintErrorReadingInput()
		return ""
	}

	if strings.ToUpper(strings.TrimSpace(wantsSSH)) != yes {
		fmt.Println("⏭️  Skipping SSH setup for now. You can configure this later by running setup again.")
		return ""
	}

	// Discover existing SSH keys
	keys, err := sshDiscovery.DiscoverSSHKeys()
	if err != nil {
		fmt.Printf("⚠️  Could not scan for SSH keys: %v\n", err)
		return s.manualSSHKeyInput()
	}

	// Show existing keys
	existingKeys := make([]SSHKeyInfo, 0)
	for _, key := range keys {
		if key.Exists {
			existingKeys = append(existingKeys, key)
		}
	}

	if len(existingKeys) > 0 {
		fmt.Println("\n📋 Found existing SSH keys:")
		for i, key := range existingKeys {
			status := "❌"
			if key.HasPublic {
				status = "✅"
			}
			fmt.Printf("   %d. %s %s (%s) %s\n",
				i+1,
				status,
				filepath.Base(key.Path),
				key.Type,
				key.Description)
		}
		fmt.Println("   ✅ = Has public key  ❌ = Missing public key")
	}

	fmt.Println("\nWhat would you like to do?")
	options := []string{
		"📁 Use existing SSH key",
		"🆕 Generate new SSH key",
		"📝 Enter SSH key path manually",
		"❌ Skip SSH setup for now",
		"❓ Show SSH setup guide",
	}

	for i, option := range options {
		fmt.Printf("   %d. %s\n", i+1, option)
	}

	choice := s.getChoice(len(options))

	switch choice {
	case 1: // Use existing key
		return s.selectExistingSSHKey(existingKeys, sshDiscovery)
	case 2: // Generate new key
		return s.generateNewSSHKey(mode, email, sshDiscovery)
	case 3: // Manual input
		return s.manualSSHKeyInput()
	case 4: // Skip
		fmt.Println("⏭️  Skipping SSH setup. You can configure this later.")
		return ""
	case 5: // Show guide
		sshDiscovery.ShowSSHSetupGuide()
		return s.setupSSHKeyForAccount(mode, email, sshDiscovery) // Recursive call
	default:
		return ""
	}
}

func (s *SetupService) selectExistingSSHKey(keys []SSHKeyInfo, sshDiscovery ISSHDiscoveryService) string {
	if len(keys) == 0 {
		fmt.Println("No existing SSH keys found.")
		return ""
	}

	fmt.Println("\nSelect an SSH key:")
	for i, key := range keys {
		fmt.Printf("   %d. %s (%s)\n", i+1, filepath.Base(key.Path), key.Type)
		if key.HasPublic {
			fmt.Printf("      Public key available ✅\n")
		} else {
			fmt.Printf("      ⚠️  Public key missing - you'll need to add it to GitHub/GitLab manually\n")
		}
	}

	choice := s.getChoice(len(keys))
	selectedKey := keys[choice-1]

	// Validate the selected key
	keyInfo, err := sshDiscovery.ValidateAndShowKeyInfo(selectedKey.Path)
	if err != nil {
		fmt.Printf("❌ Error with selected key: %v\n", err)
		fmt.Println("You can fix this later and update your configuration.")
		return selectedKey.Path // Still return it for now
	}

	fmt.Printf("✅ Selected: %s (%s)\n", filepath.Base(selectedKey.Path), keyInfo.Type)

	// Show public key if available
	if keyInfo.HasPublic {
		fmt.Println("\n📋 Your public key:")
		if content, err := sshDiscovery.GetPublicKeyContent(selectedKey.Path); err == nil {
			fmt.Printf("   %s\n", content)
		}

		publicKeyPath := selectedKey.Path
		if !strings.HasSuffix(selectedKey.Path, ".pub") {
			publicKeyPath = selectedKey.Path + ".pub"
		}

		sshDiscovery.ShowGitHubSetupGuide(publicKeyPath)
	}

	return selectedKey.Path
}

func (s *SetupService) generateNewSSHKey(mode, email string, sshDiscovery ISSHDiscoveryService) string {
	fmt.Println("\n🆕 Generate New SSH Key")
	fmt.Println("========================")

	// Choose key type
	fmt.Println("Choose SSH key type:")
	keyTypes := []string{
		"Ed25519 (recommended - modern, secure, fast)",
		"RSA 4096 (widely compatible, larger keys)",
	}

	for i, keyType := range keyTypes {
		fmt.Printf("   %d. %s\n", i+1, keyType)
	}

	typeChoice := s.getChoice(len(keyTypes))

	var keyType, filename string
	switch typeChoice {
	case 1:
		keyType = "ed25519"
		filename = fmt.Sprintf("id_ed25519_%s", mode)
	case 2:
		keyType = "rsa"
		filename = fmt.Sprintf("id_rsa_%s", mode)
	}

	fmt.Printf("\n💡 Suggested filename: %s\n", filename)
	fmt.Print("Press Enter to use this name, or type a different name: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input != "" {
			filename = input
		}
	}

	// Generate the key
	if err := sshDiscovery.GenerateSSHKey(email, keyType, filename); err != nil {
		fmt.Printf("❌ Failed to generate SSH key: %v\n", err)
		fmt.Println("You can create SSH keys manually later and update your configuration.")
		return ""
	}

	homeDir, _ := os.UserHomeDir()
	keyPath := filepath.Join(homeDir, ".ssh", filename)

	// Show next steps
	fmt.Println("\n🎉 SSH key generated successfully!")
	sshDiscovery.ShowGitHubSetupGuide(keyPath + ".pub")

	return keyPath
}

func (s *SetupService) manualSSHKeyInput() string {
	fmt.Println("\n📝 Manual SSH Key Path Entry")
	fmt.Println("==============================")

	homeDir, _ := os.UserHomeDir()
	defaultPath := filepath.Join(homeDir, ".ssh", "id_ed25519")

	fmt.Printf("Enter the full path to your SSH private key:\n")
	fmt.Printf("(default: %s): ", defaultPath)

	scanner := bufio.NewScanner(os.Stdin)
	var keyPath string
	if scanner.Scan() {
		keyPath = strings.TrimSpace(scanner.Text())
		if keyPath == "" {
			keyPath = defaultPath
		}
	}

	// Expand tilde
	if strings.HasPrefix(keyPath, "~/") {
		keyPath = filepath.Join(homeDir, keyPath[2:])
	}

	// Validate the path
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		fmt.Printf("⚠️  File does not exist: %s\n", keyPath)
		fmt.Println("You can create this key later and update your configuration.")
		return keyPath // Still return it - user might create it later
	}

	fmt.Printf("✅ SSH key path set: %s\n", keyPath)
	return keyPath
}

func (s *SetupService) getChoice(maxChoice int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter your choice (1-%d): ", maxChoice)
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if choice, err := strconv.Atoi(input); err == nil && choice >= 1 && choice <= maxChoice {
				return choice
			}
		}
		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

func (s *SetupService) checkForEmptyAccountData(savedAccounts *models.Accounts) {
	if inputPersonalEmail == "" || inputPersonalUsername == "" {
		inputPersonalEmail = savedAccounts.Personal.Email
		inputPersonalUsername = savedAccounts.Personal.Username
		if inputPersonalSigningKeyID == "" {
			inputPersonalSigningKeyID = savedAccounts.Personal.SigningKeyID
		}
		if inputPersonalSSHKeyPath == "" {
			inputPersonalSSHKeyPath = savedAccounts.Personal.SSHKeyPath
		}
	}

	if inputWorkEmail == "" || inputWorkUsername == "" {
		inputWorkEmail = savedAccounts.Work.Email
		inputWorkUsername = savedAccounts.Work.Username
		if inputWorkSigningKeyID == "" {
			inputWorkSigningKeyID = savedAccounts.Work.SigningKeyID
		}
		if inputWorkSSHKeyPath == "" {
			inputWorkSSHKeyPath = savedAccounts.Work.SSHKeyPath
		}
	}

	if inputSchoolEmail == "" || inputSchoolUsername == "" {
		inputSchoolEmail = savedAccounts.School.Email
		inputSchoolUsername = savedAccounts.School.Username
		if inputSchoolSigningKeyID == "" {
			inputSchoolSigningKeyID = savedAccounts.School.SigningKeyID
		}
		if inputSchoolSSHKeyPath == "" {
			inputSchoolSSHKeyPath = savedAccounts.School.SSHKeyPath
		}
	}
}
