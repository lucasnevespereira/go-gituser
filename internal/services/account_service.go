package services

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/models"
	"go-gituser/internal/storage"
	"strings"
)

type IAccountService interface {
	Switch(mode string) error
	ReadSavedAccounts() (*models.Accounts, error)
	ReadCurrentGitAccount() *models.Account
	CheckSavedAccount(account *models.Account) (bool, error)
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
		if savedAccounts.Work.Username == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(&savedAccounts.Work)
	case models.SchoolMode:
		if savedAccounts.School.Username == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(&savedAccounts.School)
	case models.PersonalMode:
		if savedAccounts.Personal.Username == "" {
			return models.ErrNoAccountFound
		}
		s.git.SetConfig(&savedAccounts.Personal)
	}

	return nil
}

func (s *AccountService) ReadSavedAccounts() (*models.Accounts, error) {
	savedAccounts, err := s.storage.GetAccounts()
	if err != nil {
		return nil, err
	}

	return savedAccounts, nil
}

func (s *AccountService) ReadCurrentGitAccount() *models.Account {
	currGitAccount := s.git.ReadConfig()
	currGitAccount.Username = strings.TrimSuffix(currGitAccount.Username, "\n")
	currGitAccount.Email = strings.TrimSuffix(currGitAccount.Email, "\n")
	return currGitAccount
}

func (s *AccountService) CheckSavedAccount(account *models.Account) (bool, error) {
	savedAccounts, err := s.ReadSavedAccounts()
	if err != nil {
		return false, err
	}
	return usernameIsSaved(savedAccounts, account.Username) && emailIsSaved(savedAccounts, account.Email), nil
}

func (s *AccountService) SaveAccounts(accounts *models.Accounts) error {
	err := s.storage.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	return nil
}

func usernameIsSaved(savedAccounts *models.Accounts, username string) bool {
	return savedAccounts.Personal.Username == username ||
		savedAccounts.Work.Username == username || savedAccounts.School.Username == username
}

func emailIsSaved(savedAccounts *models.Accounts, email string) bool {
	return savedAccounts.Personal.Email == email ||
		savedAccounts.Work.Email == email || savedAccounts.School.Email == email
}
