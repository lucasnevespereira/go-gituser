package services

import (
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"strings"
)

type IAccountService interface {
	Switch(mode string) error
	GetSavedAccounts() (*models.Accounts, error)
	GetCurrentGitAccount() *models.Account
	CheckSavedAccount(account *models.Account) (bool, error)
	SaveAccounts(accounts *models.Accounts) error
	SwitchSSHKey(account *models.Account) error
	ClearAllSSHKeys() error
}

type AccountService struct {
	storage storage.IAccountJSONStorage
	git     git.IConnector
	ssh     ssh.ISSHConnector
}

func NewAccountService(accountStorage storage.IAccountJSONStorage, gitConnector git.IConnector, sshConnector ssh.ISSHConnector) IAccountService {
	return &AccountService{storage: accountStorage, git: gitConnector, ssh: sshConnector}
}

func (s *AccountService) Switch(mode string) error {
	savedAccounts, err := s.GetSavedAccounts()
	if err != nil {
		return models.ErrNoAccountFound
	}
	targetAccount, ok := savedAccounts.Get(mode)
	if !ok {
		return models.ErrNoAccountFound
	}

	// Clear all SSH keys from agent before switching
	if err := s.ClearAllSSHKeys(); err != nil {
		// Don't fail the entire operation if SSH clearing fails
		fmt.Printf("⚠️ Warning: Could not clear SSH keys: %v\n", err)
	}

	// Set git configuration
	s.git.SetConfig(&targetAccount)

	// Set SSH key
	if err := s.SwitchSSHKey(&targetAccount); err != nil {
		fmt.Printf("⚠️  Warning: Could not configure SSH key: %v\n", err)
	}

	return nil
}

func (s *AccountService) GetSavedAccounts() (*models.Accounts, error) {
	savedAccounts, err := s.storage.GetAccounts()
	if err != nil {
		return nil, err
	}

	return savedAccounts, nil
}

func (s *AccountService) GetCurrentGitAccount() *models.Account {
	currGitAccount := s.git.ReadConfig()
	currGitAccount.Username = strings.TrimSuffix(currGitAccount.Username, "\n")
	currGitAccount.Email = strings.TrimSuffix(currGitAccount.Email, "\n")
	currGitAccount.SigningKeyID = strings.TrimSuffix(currGitAccount.SigningKeyID, "\n")

	foundAccount, _ := s.storage.GetAccountByUsername(currGitAccount.Username)
	if foundAccount != nil && foundAccount.SSHKeyPath != "" {
		if loaded := s.ssh.IsKeyLoaded(foundAccount.SSHKeyPath + ".pub"); !loaded {
			currGitAccount.SSHKeyPath = ""
		} else {
			currGitAccount.SSHKeyPath = foundAccount.SSHKeyPath
		}
	}

	return currGitAccount
}

func (s *AccountService) CheckSavedAccount(account *models.Account) (bool, error) {
	savedAccounts, err := s.GetSavedAccounts()
	if err != nil {
		fmt.Println("Error getting saved accounts:", err)
		return false, err
	}

	found := false
	savedAccounts.ForEachConfigured(func(_ string, saved models.Account) bool {
		if saved.Username == account.Username && saved.Email == account.Email {
			found = true
			return false
		}
		return true
	})
	return found, nil
}

func (s *AccountService) SaveAccounts(accounts *models.Accounts) error {
	err := s.storage.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) SwitchSSHKey(account *models.Account) error {
	if account.SSHKeyPath == "" {
		fmt.Println("ℹ️ No SSH key configured for this account")
		return nil
	}

	return s.ssh.AddKeyToAgent(account.SSHKeyPath)
}

func (s *AccountService) ClearAllSSHKeys() error {
	return s.ssh.ClearAgent()
}
