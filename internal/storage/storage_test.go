package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
)

func TestLegacyAccountsFileKeepsCustomModesOnSave(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	configDir := filepath.Join(home, ".config", "gituser")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatal(err)
	}
	legacy := []byte(`{"personal":{"username":"alice","email":"alice@example.com"},"work":{"username":"bob","email":"bob@example.com"},"school":{}}`)
	if err := os.WriteFile(filepath.Join(configDir, AccountsStorageFile), legacy, 0600); err != nil {
		t.Fatal(err)
	}
	storage := NewAccountJSONStorage(AccountsStorageFile)
	accounts, err := storage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	accounts.Set("freelance", models.Account{Username: "carol", Email: "carol@example.com"})
	if err := storage.SaveAccounts(accounts); err != nil {
		t.Fatal(err)
	}
	got, err := storage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if got.Personal.Username != "alice" || got.Work.Username != "bob" {
		t.Fatalf("built-in accounts were lost: %+v", got)
	}
	if account, ok := got.Get("freelance"); !ok || account.Username != "carol" {
		t.Fatalf("custom account was lost: %+v", got.Custom)
	}
	account, err := storage.GetAccountByUsername("carol")
	if err != nil || account.Email != "carol@example.com" {
		t.Fatalf("lookup by custom username = %+v, %v", account, err)
	}
}
