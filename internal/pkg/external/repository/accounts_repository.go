package repository

import (
	"go-gituser/internal/pkg/models"
	"gorm.io/gorm"
	"log"
)

type repository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	err := db.AutoMigrate(&models.Accounts{})
	if err != nil {
		log.Printf("failed to migrate accounts table: %v", err)
	}
	return &repository{
		db: db,
	}
}

func (r *repository) SaveAccounts(account models.Accounts) error {
	err := r.db.Save(&account).Error
	if err != nil {
		log.Printf("failed to save accounts: %v", err)
		return err
	}
	return nil
}

func (r *repository) GetAccounts() (models.Accounts, error) {
	var accounts models.Accounts
	err := r.db.First(&accounts).Error
	if err != nil {
		log.Printf("failed to get accounts: %v", err)
		return models.Accounts{}, err
	}
	return accounts, nil
}
