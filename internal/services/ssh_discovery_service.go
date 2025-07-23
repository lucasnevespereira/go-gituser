package services

import (
	"bufio"
	"fmt"
	"go-gituser/internal/connectors/ssh"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type SSHKeyInfo struct {
	Path        string
	Type        string
	Description string
	Exists      bool
	HasPublic   bool
	CreatedTime time.Time
}

type ISSHDiscoveryService interface {
	DiscoverSSHKeys() ([]SSHKeyInfo, error)
	GenerateSSHKey(email, keyType, filename string) error
	ShowSSHSetupGuide()
	ShowGitHubSetupGuide(publicKeyPath string)
	GetPublicKeyContent(privateKeyPath string) (string, error)
	ValidateAndShowKeyInfo(keyPath string) (*SSHKeyInfo, error)
}

type SSHDiscoveryService struct {
	sshConnector ssh.ISSHConnector
}

func NewSSHDiscoveryService(sshConnector ssh.ISSHConnector) ISSHDiscoveryService {
	return &SSHDiscoveryService{
		sshConnector: sshConnector,
	}
}

func (s *SSHDiscoveryService) DiscoverSSHKeys() ([]SSHKeyInfo, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	var keys []SSHKeyInfo

	// Common SSH key patterns
	keyPatterns := map[string]string{
		"id_rsa":     "RSA key (most common, good compatibility)",
		"id_ed25519": "Ed25519 key (modern, recommended for new keys)",
		"id_ecdsa":   "ECDSA key (elliptic curve, good security)",
		"id_dsa":     "DSA key (legacy, not recommended)",
	}

	for keyName, description := range keyPatterns {
		keyPath := filepath.Join(sshDir, keyName)
		publicKeyPath := keyPath + ".pub"

		keyInfo := SSHKeyInfo{
			Path:        keyPath,
			Type:        strings.ToUpper(strings.TrimPrefix(keyName, "id_")),
			Description: description,
			Exists:      false,
			HasPublic:   false,
		}

		// Check if private key exists
		if stat, err := os.Stat(keyPath); err == nil {
			keyInfo.Exists = true
			keyInfo.CreatedTime = stat.ModTime()
		}

		// Check if public key exists
		if _, err := os.Stat(publicKeyPath); err == nil {
			keyInfo.HasPublic = true
		}

		keys = append(keys, keyInfo)
	}

	// Look for other SSH keys in the directory
	if entries, err := os.ReadDir(sshDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := entry.Name()
			// Skip known keys and public keys
			if strings.HasSuffix(name, ".pub") ||
				name == "id_rsa" || name == "id_ed25519" ||
				name == "id_ecdsa" || name == "id_dsa" ||
				name == "known_hosts" || name == "config" {
				continue
			}

			// Check if it looks like a private key
			keyPath := filepath.Join(sshDir, name)
			if s.looksLikePrivateKey(keyPath) {
				info, _ := entry.Info()
				keyInfo := SSHKeyInfo{
					Path:        keyPath,
					Type:        "CUSTOM",
					Description: "Custom SSH key",
					Exists:      true,
					HasPublic:   false,
					CreatedTime: info.ModTime(),
				}

				// Check for corresponding public key
				if _, err := os.Stat(keyPath + ".pub"); err == nil {
					keyInfo.HasPublic = true
				}

				keys = append(keys, keyInfo)
			}
		}
	}

	// Sort keys by existence and then by type preference
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Exists != keys[j].Exists {
			return keys[i].Exists // Existing keys first
		}
		// Prefer ed25519, then rsa, then others
		preference := map[string]int{
			"ED25519": 0,
			"RSA":     1,
			"ECDSA":   2,
			"CUSTOM":  3,
			"DSA":     4,
		}
		return preference[keys[i].Type] < preference[keys[j].Type]
	})

	return keys, nil
}

func (s *SSHDiscoveryService) looksLikePrivateKey(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return strings.Contains(firstLine, "PRIVATE KEY") ||
			strings.Contains(firstLine, "BEGIN OPENSSH PRIVATE KEY") ||
			strings.Contains(firstLine, "BEGIN RSA PRIVATE KEY") ||
			strings.Contains(firstLine, "BEGIN EC PRIVATE KEY")
	}
	return false
}

func (s *SSHDiscoveryService) GenerateSSHKey(email, keyType, filename string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	keyPath := filepath.Join(sshDir, filename)

	fmt.Printf("🔧 Generating %s SSH key...\n", keyType)

	var cmd []string
	switch strings.ToLower(keyType) {
	case "ed25519":
		cmd = []string{"ssh-keygen", "-t", "ed25519", "-C", email, "-f", keyPath}
	case "rsa":
		cmd = []string{"ssh-keygen", "-t", "rsa", "-b", "4096", "-C", email, "-f", keyPath}
	default:
		return fmt.Errorf("unsupported key type: %s", keyType)
	}

	fmt.Println("💡 You'll be prompted for a passphrase. You can:")
	fmt.Println("   - Press Enter for no passphrase (less secure but convenient)")
	fmt.Println("   - Enter a passphrase for additional security")
	fmt.Println()

	// Execute the command interactively
	execCmd := exec.Command(cmd[0], cmd[1:]...)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	fmt.Printf("✅ SSH key generated successfully!\n")
	fmt.Printf("   Private key: %s\n", keyPath)
	fmt.Printf("   Public key:  %s.pub\n", keyPath)

	return nil
}

//	func (s *SSHDiscoveryService) ShowSSHSetupGuide() {
//		fmt.Println()
//		fmt.Println("🔑 SSH Key Setup Guide")
//		fmt.Println("======================")
//		fmt.Println()
//		fmt.Println("SSH keys allow you to authenticate with Git repositories securely.")
//		fmt.Println("You'll need separate SSH keys for each account (work, personal, school).")
//		fmt.Println()
//		fmt.Println("📍 Where SSH keys are stored:")
//		homeDir, _ := os.UserHomeDir()
//		fmt.Printf("   %s/.ssh/\n", homeDir)
//		fmt.Println()
//		fmt.Println("🔍 Common SSH key files:")
//		fmt.Println("   id_ed25519     - Private key (recommended for new keys)")
//		fmt.Println("   id_ed25519.pub - Public key (share this with GitHub/GitLab)")
//		fmt.Println("   id_rsa         - Private RSA key (widely compatible)")
//		fmt.Println("   id_rsa.pub     - Public RSA key")
//		fmt.Println()
//		fmt.Println("💡 Pro tip: Name your keys descriptively!")
//		fmt.Println("   Examples: id_ed25519_work, id_rsa_personal, id_ed25519_school")
//		fmt.Println()
//	}

func (s *SSHDiscoveryService) ShowSSHSetupGuide() {
	fmt.Println()
	fmt.Println("🔑 SSH Key Setup Guide")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println()

	fmt.Println("📖 WHAT ARE SSH KEYS?")
	fmt.Println("─────────────────────")
	fmt.Println("SSH keys allow you to authenticate with Git repositories securely without")
	fmt.Println("entering your password every time. They consist of two parts:")
	fmt.Println("  • Private key (keep secret) - stays on your computer")
	fmt.Println("  • Public key (share freely) - upload to GitHub/GitLab")
	fmt.Println()

	fmt.Println("🏗️  GITUSER WORKFLOW")
	fmt.Println("───────────────────")
	fmt.Println("GitUser helps you manage separate SSH keys for each account:")
	fmt.Println("  • Work account    → work SSH key")
	fmt.Println("  • Personal account → personal SSH key")
	fmt.Println("  • School account  → school SSH key")
	fmt.Println()
	fmt.Println("When you switch accounts, GitUser automatically loads the correct SSH key!")
	fmt.Println()

	fmt.Println("📍 WHERE SSH KEYS LIVE")
	fmt.Println("─────────────────────")
	homeDir, _ := os.UserHomeDir()
	fmt.Printf("All SSH keys are stored in: %s/.ssh/\n", homeDir)
	fmt.Println()

	fmt.Println("Common SSH key files:")
	fmt.Println("┌──────────────────┬─────────────────────────────────────────────┐")
	fmt.Println("│ File             │ Description                                 │")
	fmt.Println("├──────────────────┼─────────────────────────────────────────────┤")
	fmt.Println("│ id_ed25519       │ Private Ed25519 key (recommended)          │")
	fmt.Println("│ id_ed25519.pub   │ Public Ed25519 key (upload to GitHub)      │")
	fmt.Println("│ id_rsa           │ Private RSA key (widely compatible)        │")
	fmt.Println("│ id_rsa.pub       │ Public RSA key (upload to GitHub)          │")
	fmt.Println("└──────────────────┴─────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("💡 NAMING RECOMMENDATIONS")
	fmt.Println("─────────────────────────")
	fmt.Println("Use descriptive names for multiple accounts:")
	fmt.Println("  • id_ed25519_work      ← for work projects")
	fmt.Println("  • id_ed25519_personal  ← for personal projects")
	fmt.Println("  • id_rsa_school        ← for school projects")
	fmt.Println()

	fmt.Println("🚀 GETTING STARTED")
	fmt.Println("─────────────────")
	fmt.Println("1. Check what you have:    gituser ssh discover")
	fmt.Println("2. Set up your accounts:   gituser setup")
	fmt.Println("3. Test connections:       gituser ssh test")
	fmt.Println("4. Switch between accounts: gituser work")
	fmt.Println()

	fmt.Println("🆘 NEED HELP?")
	fmt.Println("────────────")
	fmt.Println("  • gituser ssh discover  - Find existing keys")
	fmt.Println("  • gituser ssh test      - Test GitHub/GitLab connections")
	fmt.Println("  • gituser setup         - Interactive setup wizard")
	fmt.Println()
}

// func (s *SSHDiscoveryService) ShowGitHubSetupGuide(publicKeyPath string) {
// 	fmt.Println()
// 	fmt.Println("🐙 Adding SSH Key to GitHub/GitLab")
// 	fmt.Println("===================================")
// 	fmt.Println()

// 	if publicKeyPath != "" {
// 		fmt.Println("1. Copy your public key:")
// 		content, err := s.GetPublicKeyContent(publicKeyPath)
// 		if err == nil {
// 			fmt.Printf("   %s\n", content)
// 		} else {
// 			fmt.Printf("   cat %s\n", publicKeyPath)
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println("2. Add to GitHub:")
// 	fmt.Println("   • Go to GitHub.com → Settings → SSH and GPG keys")
// 	fmt.Println("   • Click 'New SSH key'")
// 	fmt.Println("   • Give it a descriptive title (e.g., 'Work Laptop')")
// 	fmt.Println("   • Paste your public key")
// 	fmt.Println()
// 	fmt.Println("3. Add to GitLab:")
// 	fmt.Println("   • Go to GitLab.com → Preferences → SSH Keys")
// 	fmt.Println("   • Paste your public key")
// 	fmt.Println("   • Add a descriptive title")
// 	fmt.Println()
// 	fmt.Println("4. Test your connection:")
// 	fmt.Println("   ssh -T git@github.com")
// 	fmt.Println("   ssh -T git@gitlab.com")
// 	fmt.Println()
// }

func (s *SSHDiscoveryService) ShowGitHubSetupGuide(publicKeyPath string) {
	fmt.Println()
	fmt.Println("🐙 Adding SSH Key to GitHub/GitLab")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println()

	if publicKeyPath != "" {
		fmt.Println("📋 STEP 1: COPY YOUR PUBLIC KEY")
		fmt.Println("──────────────────────────────────")
		fmt.Println("Your public key (copy this entire text):")
		fmt.Println()

		content, err := s.GetPublicKeyContent(publicKeyPath)
		if err == nil {
			fmt.Printf("┌─────────────────────────────────────────────────────────────┐\n")
			fmt.Printf("│ %-59s │\n", content[:min(59, len(content))])
			if len(content) > 59 {
				// Handle long keys by wrapping
				remaining := content[59:]
				for len(remaining) > 0 {
					line := remaining[:min(59, len(remaining))]
					fmt.Printf("│ %-59s │\n", line)
					if len(remaining) <= 59 {
						break
					}
					remaining = remaining[59:]
				}
			}
			fmt.Printf("└─────────────────────────────────────────────────────────────┘\n")
		} else {
			fmt.Printf("💻 Or copy with command: cat %s\n", publicKeyPath)
		}
		fmt.Println()
	}

	fmt.Println("🐙 STEP 2: ADD TO GITHUB")
	fmt.Println("────────────────────────")
	fmt.Println("1. Go to: https://github.com/settings/keys")
	fmt.Println("2. Click the green 'New SSH key' button")
	fmt.Println("3. Fill in the form:")
	fmt.Println("   • Title: Give it a name (e.g., 'Work Laptop - GitUser')")
	fmt.Println("   • Key type: Authentication Key")
	fmt.Println("   • Key: Paste your public key from above")
	fmt.Println("4. Click 'Add SSH key'")
	fmt.Println()

	fmt.Println("🦊 STEP 3: ADD TO GITLAB")
	fmt.Println("────────────────────────")
	fmt.Println("1. Go to: https://gitlab.com/-/profile/keys")
	fmt.Println("2. Fill in the form:")
	fmt.Println("   • Title: Give it a name (e.g., 'Work Laptop - GitUser')")
	fmt.Println("   • Key: Paste your public key from above")
	fmt.Println("   • Usage type: Authentication & Signing (recommended)")
	fmt.Println("3. Click 'Add key'")
	fmt.Println()

	fmt.Println("🧪 STEP 4: TEST YOUR CONNECTION")
	fmt.Println("──────────────────────────────")
	fmt.Println("Run these commands to test:")
	fmt.Println("  gituser ssh test")
	fmt.Println("  ssh -T git@github.com")
	fmt.Println("  ssh -T git@gitlab.com")
	fmt.Println()

	fmt.Println("✅ SUCCESS MESSAGES TO LOOK FOR:")
	fmt.Println("GitHub: 'Hi username! You've successfully authenticated...'")
	fmt.Println("GitLab: 'Welcome to GitLab, @username!'")
	fmt.Println()
}

// Helper function for string length
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *SSHDiscoveryService) GetPublicKeyContent(privateKeyPath string) (string, error) {
	publicKeyPath := privateKeyPath + ".pub"
	content, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func (s *SSHDiscoveryService) ValidateAndShowKeyInfo(keyPath string) (*SSHKeyInfo, error) {
	if err := s.sshConnector.ValidateKeyPath(keyPath); err != nil {
		return nil, err
	}

	keyInfo := &SSHKeyInfo{
		Path:   keyPath,
		Exists: true,
	}

	// Try to determine key type
	content, err := os.ReadFile(keyPath)
	if err != nil {
		return keyInfo, nil
	}

	contentStr := string(content)
	switch {
	case strings.Contains(contentStr, "BEGIN OPENSSH PRIVATE KEY"):
		if strings.Contains(contentStr, "ed25519") {
			keyInfo.Type = "ED25519"
		} else {
			keyInfo.Type = "OPENSSH"
		}
	case strings.Contains(contentStr, "BEGIN RSA PRIVATE KEY"):
		keyInfo.Type = "RSA"
	case strings.Contains(contentStr, "BEGIN EC PRIVATE KEY"):
		keyInfo.Type = "ECDSA"
	default:
		keyInfo.Type = "UNKNOWN"
	}

	// Check if public key exists
	if _, err := os.Stat(keyPath + ".pub"); err == nil {
		keyInfo.HasPublic = true
	}

	return keyInfo, nil
}
