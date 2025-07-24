package storage

import (
	"encoding/json"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"os"

	"github.com/pkg/errors"
)

const (
	AccountsStorageFile = "accounts.json"
	gituserConfigPath   = "/.config/gituser/"
)

type IAccountJSONStorage interface {
	GetAccounts() (*models.Accounts, error)
	GetAccountByUsername(username string) (*models.Account, error)
	SaveAccounts(accounts *models.Accounts) error
}

type AccountJSONStorage struct {
	file string
}

func NewAccountJSONStorage(file string) IAccountJSONStorage {
	return &AccountJSONStorage{file: file}
}

func (s *AccountJSONStorage) GetAccounts() (*models.Accounts, error) {
	accountsFile, err := s.readAccountsFile()
	if err != nil {
		return nil, err
	}

	data, err := readFileData(accountsFile)
	if err != nil {
		return nil, err
	}

	var rowAccounts *models.Accounts
	err = json.Unmarshal(data, &rowAccounts)
	if err != nil {
		return nil, err
	}

	return rowAccounts, nil
}

func (s *AccountJSONStorage) GetAccountByUsername(username string) (*models.Account, error) {
	accounts, err := s.GetAccounts()
	if err != nil {
		return nil, err
	}

	if accounts.Personal.Username == username {
		return &accounts.Personal, nil
	}
	if accounts.Work.Username == username {
		return &accounts.Work, nil
	}
	if accounts.School.Username == username {
		return &accounts.School, nil
	}

	return nil, errors.New("account not found")
}

func (s *AccountJSONStorage) SaveAccounts(accounts *models.Accounts) error {
	accountsFile, err := s.readAccountsFile()
	if err != nil {
		return errors.Wrap(err, "storage.WriteAccountsData.GetAccountsDataFile")
	}

	file, err := json.MarshalIndent(accounts, "", " ")
	if err != nil {
		logger.PrintErrorWithMessage(err, "json.MarshalIndent")
		return errors.Wrap(err, "storage.WriteAccountsData.json.MarshalIndent")
	}

	err = os.WriteFile(accountsFile, file, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "storage.WriteAccountsData.writeFile")
	}

	return nil
}

func (s *AccountJSONStorage) readAccountsFile() (string, error) {
	dataFile, err := getLocalFile(s.file)
	if err != nil {
		return "", errors.Wrap(err, "storage.readAccountsFile.getLocalFile")
	}

	return dataFile, nil
}

func ensureLocalConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "storage.ensureLocalConfigDir.UserHomeDir")
	}
	configDir := homeDir + gituserConfigPath
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return "", errors.Wrap(err, "storage.ensureLocalConfigDir.MkdirAll")
		}
	}

	return configDir, nil
}

func getLocalFile(filename string) (string, error) {
	localConfigDir, err := ensureLocalConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "storage.getLocalFile.ensureLocalConfigDir")
	}

	localDataFile := localConfigDir + filename
	_, err = os.Stat(localDataFile)
	if os.IsNotExist(err) {
		createdFile, err := os.Create(localDataFile)
		if err != nil {
			return "", errors.Wrap(err, "storage.getLocalFile.create")
		}

		_, err = createdFile.WriteString("{}")
		if err != nil {
			return "", errors.Wrap(err, "storage.getLocalFile.write")
		}
	}
	return localDataFile, nil
}

func readFileData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, errors.Wrap(err, "storage.ReadFileData")
	}
	return data, nil
}
