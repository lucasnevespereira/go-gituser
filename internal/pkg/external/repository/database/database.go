package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const AccountsDBPath = "accounts.db"

func Open(dbPath string) (*gorm.DB, error) {
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create the database file if it doesn't exist
		_, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
	}

	// Open the database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
