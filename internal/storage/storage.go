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
	ReadAccountsFile() (string, error)
	WriteAccountsFile(accounts *models.Accounts) error
}

type AccountJSONStorage struct {
	file string
}

func NewAccountJSONStorage(file string) IAccountJSONStorage {
	return &AccountJSONStorage{file: file}
}

func (s *AccountJSONStorage) ReadAccountsFile() (string, error) {
	dataFile, err := getLocalFile(s.file)
	if err != nil {
		return "", errors.Wrap(err, "storage.GetAccountsDataFile")
	}

	return dataFile, nil
}

func (s *AccountJSONStorage) WriteAccountsFile(accounts *models.Accounts) error {
	accountsFile, err := s.ReadAccountsFile()
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
