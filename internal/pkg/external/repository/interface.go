package repository

import "go-gituser/internal/pkg/models"

type AccountRepository interface {
	SaveAccounts(account models.Accounts) error
	GetAccounts() (models.Accounts, error)
}
