package services

import (
	"encoding/json"
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/storage"
	"strings"
)

type IAccountService interface {
	Switch(mode string) error
	ReadSavedAccounts() (*models.Accounts, error)
	ReadCurrentGitAccount() (name, email string)
	CheckSavedAccount(username, email string) (bool, error)
	SaveAccounts(accounts *models.Accounts) error
}

type AccountService struct {
	git     git.IConnector
	storage storage.IAccountJSONStorage
}

func NewAccountService(accountStorage storage.IAccountJSONStorage, gitConnector git.IConnector) IAccountService {
	return &AccountService{storage: accountStorage, git: gitConnector}
}

func (s *AccountService) Switch(mode string) error {
	savedAccounts, err := s.ReadSavedAccounts()
	if err != nil {
		return models.ErrNoAccountFound
	}

	switch mode {
	case models.WorkMode:
		if savedAccounts.WorkUsername == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(savedAccounts.WorkUsername, savedAccounts.WorkEmail)
	case models.SchoolMode:
		if savedAccounts.SchoolUsername == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(savedAccounts.SchoolUsername, savedAccounts.SchoolEmail)
	case models.PersonalMode:
		if savedAccounts.PersonalUsername == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(savedAccounts.PersonalUsername, savedAccounts.PersonalEmail)
	}

	return nil
}

func (s *AccountService) ReadSavedAccounts() (*models.Accounts, error) {
	accountsFile, err := s.storage.ReadAccountsFile()
	if err != nil {
		return nil, err
	}

	data, err := logger.ReadFileData(accountsFile)
	if err != nil {
		return nil, err
	}

	var savedAccounts *models.Accounts

	err = json.Unmarshal(data, &savedAccounts)
	if err != nil {
		return nil, err
	}

	return savedAccounts, nil
}

func (s *AccountService) ReadCurrentGitAccount() (name, email string) {
	currEmail, currName := s.git.ReadConfig()
	currName = strings.TrimSuffix(currName, "\n")
	currEmail = strings.TrimSuffix(currEmail, "\n")
	return currName, currEmail
}

func (s *AccountService) CheckSavedAccount(username, email string) (bool, error) {
	savedAccounts, err := s.ReadSavedAccounts()
	if err != nil {
		return false, err
	}
	return usernameIsSaved(savedAccounts, username) && emailIsSaved(savedAccounts, email), nil
}

func (s *AccountService) SaveAccounts(accounts *models.Accounts) error {
	err := s.storage.WriteAccountsFile(accounts)
	if err != nil {
		return err
	}

	return nil
}

func usernameIsSaved(savedAccounts *models.Accounts, username string) bool {
	return savedAccounts.PersonalUsername == username ||
		savedAccounts.WorkUsername == username || savedAccounts.SchoolUsername == username
}

func emailIsSaved(savedAccounts *models.Accounts, email string) bool {
	return savedAccounts.PersonalEmail == email ||
		savedAccounts.WorkEmail == email || savedAccounts.SchoolEmail == email
}
