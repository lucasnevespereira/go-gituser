package ssh

import (
	"bufio"
	"fmt"
	"go-gituser/internal/logger"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

type ISSHConnector interface {
	AddKeyToAgent(keyPath string) error
	RemoveKeyFromAgent(keyPath string) error
	ListKeysInAgent() ([]string, error)
	ClearAgent() error
	IsKeyLoaded(keyPath string) bool
	ValidateKeyPath(keyPath string) error

	GetDefaultKeyPath() string
	StartSSHAgent() error

	GetPublicKeyContent(publicKeyPath string) (string, error)
}

type SSHConnector struct{}

func NewSSHConnector() *SSHConnector {
	return &SSHConnector{}
}
func (s *SSHConnector) AddKeyToAgent(keyPath string) error {
	if keyPath == "" {
		return nil // no ssh key configured
	}
	if err := s.ValidateKeyPath(keyPath); err != nil {
		return fmt.Errorf("invalid SSH key path: %w", err)
	}

	// Ensure SSH agent is running
	if err := s.StartSSHAgent(); err != nil {
		return fmt.Errorf("failed to start SSH agent: %w", err)
	}

	cmd := exec.Command("ssh-add", keyPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add SSH key to agent: %s", string(output))
	}

	fmt.Printf("🔑 SSH key %s added to agent\n", filepath.Base(keyPath))
	return nil
}

func (s *SSHConnector) RemoveKeyFromAgent(keyPath string) error {
	if keyPath == "" {
		return nil
	}

	cmd := exec.Command("ssh-add", "-d", keyPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Don't treat as error if key wasn't in agent
		if strings.Contains(string(output), "not found") {
			return nil
		}
		return fmt.Errorf("failed to remove SSH key from agent: %s", string(output))
	}

	fmt.Printf("🗑️ SSH key %s removed from agent\n", filepath.Base(keyPath))
	return nil
}

func (s *SSHConnector) ListKeysInAgent() ([]string, error) {
	cmd := exec.Command("ssh-add", "-L")

	output, err := cmd.CombinedOutput()
	if err != nil {
		// If no keys are loaded, ssh-add -l returns exit code 1
		if strings.Contains(string(output), "no identities") {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to list SSH keys: %s", string(output))
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var keys []string
	for _, line := range lines {
		if line != "" {
			keys = append(keys, line)
		}
	}

	return keys, nil
}

func (s *SSHConnector) ClearAgent() error {
	cmd := exec.Command("ssh-add", "-D")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clear SSH agent: %s", string(output))
	}
	return nil
}

func (s *SSHConnector) IsKeyLoaded(keyPath string) bool {
	if keyPath == "" {
		return false
	}

	keys, err := s.ListKeysInAgent()
	if err != nil {
		return false
	}

	pubContent, err := s.GetPublicKeyContent(keyPath)
	if err != nil {
		return false
	}

	if slices.Contains(keys, pubContent) {
		return true
	}

	return false
}

func (s *SSHConnector) ValidateKeyPath(keyPath string) error {
	if keyPath == "" {
		return nil
	}

	// Expand tilde to home directory
	if strings.HasPrefix(keyPath, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get home directory: %w", err)
		}
		keyPath = filepath.Join(homeDir, keyPath[2:])
	}

	// Check if file exists
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return fmt.Errorf("SSH key file does not exist: %s", keyPath)
	}

	// Check if it's likely a private key (basic validation)
	file, err := os.Open(keyPath)
	if err != nil {
		return fmt.Errorf("cannot open SSH key file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: failed to close file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		if !strings.Contains(firstLine, "PRIVATE KEY") {
			return fmt.Errorf("file does not appear to be a private SSH key")
		}
	}

	return nil
}

func (s *SSHConnector) GetDefaultKeyPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	// Check common SSH key locations
	commonKeys := []string{
		"id_rsa",
		"id_ed25519",
		"id_ecdsa",
		"id_dsa",
	}

	for _, keyName := range commonKeys {
		keyPath := filepath.Join(homeDir, ".ssh", keyName)
		if _, err := os.Stat(keyPath); err == nil {
			return keyPath
		}
	}

	return filepath.Join(homeDir, ".ssh", "id_rsa") // Default fallback
}

func (s *SSHConnector) StartSSHAgent() error {
	// Check if SSH agent is already running
	if os.Getenv("SSH_AUTH_SOCK") != "" {
		return nil // Agent is already running
	}

	if runtime.GOOS == "windows" {
		// On Windows, try to start the ssh-agent service
		cmd := exec.Command("powershell", "-Command", "Start-Service ssh-agent")
		if err := cmd.Run(); err != nil {
			logger.PrintErrorWithMessage(err, "Failed to start SSH agent service on Windows")
			return err
		}
	} else {
		// On Unix-like systems, ssh-agent should be started by the shell/session
		// We can't easily start it from here as it needs to set environment variables
		// Let the user know they might need to start it manually
		fmt.Println("⚠️  SSH agent might not be running. You may need to run: eval $(ssh-agent)")
	}

	return nil
}

func (s *SSHConnector) GetPublicKeyContent(publicKeyPath string) (string, error) {
	content, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
