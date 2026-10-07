package services

import (
	"bufio"
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/format"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type ISetupService interface {
	SetupAccounts() error
	DeleteCustomMode(mode string) error
}

type SetupService struct {
	accountService IAccountService
	selectOption   func(label string, items []string) (int, error)
	readLine       func() (string, error)
}

func NewSetupService(accountService IAccountService) ISetupService {
	return &SetupService{
		accountService: accountService,
		selectOption: func(label string, items []string) (int, error) {
			prompt := promptui.Select{Label: label, Items: items}
			index, _, err := prompt.Run()
			return index, err
		},
		readLine: func() (string, error) {
			line, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil && line == "" {
				return "", err
			}
			return strings.TrimSpace(line), nil
		},
	}
}

const (
	workSelectLabel     = "💻 Work Account"
	schoolSelectLabel   = "📚 School Account"
	personalSelectLabel = "🏠 Personal Account"
	customSelectLabel   = "➕ New custom mode"
	deleteSelectLabel   = "🗑️ Delete custom mode"
	cancelSelectLabel   = "Cancel"
	yes                 = "Y"
)

func (s *SetupService) SetupAccounts() error {
	savedAccounts, err := s.accountService.GetSavedAccounts()
	if err != nil {
		return models.ErrSetupAccounts
	}

	for {
		items := []string{workSelectLabel, schoolSelectLabel, personalSelectLabel}
		modes := []string{models.WorkMode, models.SchoolMode, models.PersonalMode}
		customModes := make([]string, 0, len(savedAccounts.Custom))
		for mode := range savedAccounts.Custom {
			customModes = append(customModes, mode)
		}
		sort.Strings(customModes)
		for _, mode := range customModes {
			items = append(items, "🔖 "+mode)
			modes = append(modes, mode)
		}
		newCustomIndex := len(items)
		items = append(items, customSelectLabel)
		deleteIndex := len(items)
		items = append(items, deleteSelectLabel)
		cancelIndex := len(items)
		items = append(items, cancelSelectLabel)
		modes = append(modes, "", "", "")

		index, err := s.selectOption("Please choose an account to configure", items)
		if err != nil {
			return models.ErrReadingInput
		}
		if index == cancelIndex {
			return nil
		}
		if index == deleteIndex {
			if err := s.chooseCustomModeToDelete(savedAccounts); err != nil {
				return err
			}
			continue
		}

		mode := modes[index]
		if index == newCustomIndex {
			for {
				fmt.Println("What should this mode be called? (e.g. freelance)")
				if mode, err = s.readLine(); err != nil {
					return models.ErrReadingInput
				}
				if err := models.ValidateCustomMode(mode); err != nil {
					fmt.Println(err)
					continue
				}
				if _, exists := savedAccounts.Custom[mode]; exists {
					fmt.Printf("%q already exists. Select it from the menu to update it.\n", mode)
					continue
				}
				break
			}
		}

		current, _ := savedAccounts.Get(mode)
		account, err := s.selectUserAccount(mode, current)
		if err != nil {
			return err
		}
		savedAccounts.Set(mode, account)
		if err := s.accountService.SaveAccounts(savedAccounts); err != nil {
			return models.ErrSetupAccounts
		}
		logger.PrintRemeberToActiveMode(mode)

		again, err := s.confirm("Would you like to configure another account?", false)
		if err != nil {
			return err
		}
		if !again {
			fmt.Println("Okay. Bye there!")
			return nil
		}
	}
}

func (s *SetupService) chooseCustomModeToDelete(accounts *models.Accounts) error {
	modes := make([]string, 0, len(accounts.Custom))
	for mode := range accounts.Custom {
		modes = append(modes, mode)
	}
	sort.Strings(modes)
	if len(modes) == 0 {
		fmt.Println("No custom modes to delete.")
		return nil
	}

	items := append(append([]string{}, modes...), cancelSelectLabel)
	index, err := s.selectOption("Select a custom mode to delete", items)
	if err != nil {
		return models.ErrReadingInput
	}
	if index == len(modes) {
		return nil
	}
	return s.confirmAndDeleteCustomMode(accounts, modes[index])
}

func (s *SetupService) DeleteCustomMode(mode string) error {
	accounts, err := s.accountService.GetSavedAccounts()
	if err != nil {
		return models.ErrSetupAccounts
	}
	return s.confirmAndDeleteCustomMode(accounts, mode)
}

func (s *SetupService) confirmAndDeleteCustomMode(accounts *models.Accounts, mode string) error {
	if _, exists := accounts.Custom[mode]; !exists {
		return fmt.Errorf("custom mode %q not found", mode)
	}

	index, err := s.selectOption(fmt.Sprintf("Delete %q from saved accounts?", mode), []string{"No, keep it", "Yes, delete it"})
	if err != nil {
		return models.ErrReadingInput
	}
	if index != 1 {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	updated := *accounts
	updated.Custom = make(map[string]models.Account, len(accounts.Custom)-1)
	for name, account := range accounts.Custom {
		if name != mode {
			updated.Custom[name] = account
		}
	}
	if err := s.accountService.SaveAccounts(&updated); err != nil {
		return models.ErrSetupAccounts
	}
	*accounts = updated
	fmt.Printf("Custom mode %q deleted.\n", mode)
	return nil
}

// selectUserAccount asks for the account fields. When current is already
// configured, each prompt shows the saved value and Enter keeps it.
func (s *SetupService) selectUserAccount(mode string, current models.Account) (models.Account, error) {
	sshConnector := ssh.NewSSHConnector()
	sshDiscovery := NewSSHDiscoveryService(sshConnector)

	fmt.Printf("\n=== %s Account Setup ===\n", format.TitleCase(mode))
	if current.Username != "" {
		fmt.Println("💡 Press Enter to keep the current value shown in [brackets]")
	}

	var account models.Account
	var err error
	if account.Username, err = s.ask(fmt.Sprintf("What is your %s username?", mode), current.Username); err != nil {
		return account, err
	}
	if account.Email, err = s.ask(fmt.Sprintf("What is your %s email?", mode), current.Email); err != nil {
		return account, err
	}

	fmt.Printf("\n🔑 GPG Key Setup for %s Account\n", format.TitleCase(mode))
	fmt.Println("=====================================")
	useGPG, err := s.confirm("Would you like to use GPG signing for this account?", current.SigningKeyID != "")
	if err != nil {
		return account, err
	}
	if useGPG {
		if account.SigningKeyID, err = s.ask(fmt.Sprintf("What is your %s GPG signing key ID?", mode), current.SigningKeyID); err != nil {
			return account, err
		}
	}

	if current.SSHKeyPath != "" {
		fmt.Printf("\n🔑 SSH Key Setup for %s Account\n", format.TitleCase(mode))
		fmt.Println("=====================================")
		keep, err := s.confirm(fmt.Sprintf("Keep the current SSH key (%s)?", current.SSHKeyPath), true)
		if err != nil {
			return account, err
		}
		if keep {
			account.SSHKeyPath = current.SSHKeyPath
			return account, nil
		}
	}
	account.SSHKeyPath = s.setupSSHKeyForAccount(mode, account.Email, sshDiscovery)
	return account, nil
}

// ask reads a value, falling back to current on an empty answer.
func (s *SetupService) ask(question, current string) (string, error) {
	for {
		if current != "" {
			fmt.Printf("%s [%s]\n", question, current)
		} else {
			fmt.Println(question)
		}
		input, err := s.readLine()
		if err != nil {
			return "", models.ErrReadingInput
		}
		if input != "" {
			return input, nil
		}
		if current != "" {
			return current, nil
		}
		fmt.Println("❌ This field is required.")
	}
}

// confirm reads a yes/no answer, falling back to def on an empty answer.
func (s *SetupService) confirm(question string, def bool) (bool, error) {
	hint := "(y/N)"
	if def {
		hint = "(Y/n)"
	}
	fmt.Printf("%s %s\n", question, hint)
	input, err := s.readLine()
	if err != nil {
		return false, models.ErrReadingInput
	}
	if input == "" {
		return def, nil
	}
	return strings.HasPrefix(strings.ToUpper(input), yes), nil
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
