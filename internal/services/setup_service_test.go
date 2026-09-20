package services

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
)

type setupSelection struct {
	label string
	item  string
}

func scriptedSetupSelections(t *testing.T, steps []setupSelection) func(string, []string) (int, error) {
	t.Helper()
	next := 0
	t.Cleanup(func() {
		if next != len(steps) {
			t.Errorf("used %d of %d setup selections", next, len(steps))
		}
	})
	return func(label string, items []string) (int, error) {
		t.Helper()
		if next >= len(steps) {
			t.Fatalf("unexpected setup selection %q with items %v", label, items)
		}
		step := steps[next]
		next++
		if label != step.label {
			t.Fatalf("selection label = %q, want %q", label, step.label)
		}
		for index, item := range items {
			if item == step.item {
				return index, nil
			}
		}
		t.Fatalf("selection %q missing from %v", step.item, items)
		return 0, nil
	}
}

func TestSetupDeletesOnlySelectedCustomMode(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	keyPath := filepath.Join(t.TempDir(), "freelance-key")
	if err := os.WriteFile(keyPath, []byte("private key remains"), 0600); err != nil {
		t.Fatal(err)
	}
	accounts := &models.Accounts{
		Work: models.Account{Username: "worker", Email: "work@example.com"},
	}
	accounts.Set("freelance", models.Account{Username: "freelancer", Email: "freelance@example.com", SSHKeyPath: keyPath})
	accounts.Set("open-source", models.Account{Username: "contributor", Email: "oss@example.com"})
	accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
	if err := accountStorage.SaveAccounts(accounts); err != nil {
		t.Fatal(err)
	}
	git := &testGit{}
	ssh := &testSSH{}
	setup := NewSetupService(NewAccountService(accountStorage, git, ssh)).(*SetupService)
	selectOption := scriptedSetupSelections(t, []setupSelection{
		{"Please choose an account to configure", deleteSelectLabel},
		{"Select a custom mode to delete", "freelance"},
		{"Delete \"freelance\" from saved accounts?", "Yes, delete it"},
		{"Please choose an account to configure", cancelSelectLabel},
	})
	selectionCount := 0
	setup.selectOption = func(label string, items []string) (int, error) {
		selectionCount++
		if selectionCount == 4 {
			for _, item := range items {
				if item == "🔖 freelance" {
					t.Fatal("deleted mode still appears in setup menu")
				}
			}
		}
		return selectOption(label, items)
	}
	if err := setup.SetupAccounts(); err != nil {
		t.Fatal(err)
	}

	saved, err := accountStorage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := saved.Get("freelance"); ok {
		t.Fatal("deleted mode still exists in saved accounts")
	}
	if !reflect.DeepEqual(saved.Work, accounts.Work) || !reflect.DeepEqual(saved.Custom["open-source"], accounts.Custom["open-source"]) {
		t.Fatalf("other accounts changed: %+v", saved)
	}
	if _, err := os.Stat(keyPath); err != nil {
		t.Fatalf("SSH key was removed: %v", err)
	}
	if git.configured != nil || ssh.cleared != 0 || ssh.added != "" {
		t.Fatal("deleting a saved mode changed Git configuration or SSH agent")
	}
}

func TestSetupDeletionCanBeCancelled(t *testing.T) {
	for _, tc := range []struct {
		name  string
		steps []setupSelection
	}{
		{"at mode picker", []setupSelection{
			{"Please choose an account to configure", deleteSelectLabel},
			{"Select a custom mode to delete", cancelSelectLabel},
			{"Please choose an account to configure", cancelSelectLabel},
		}},
		{"at confirmation", []setupSelection{
			{"Please choose an account to configure", deleteSelectLabel},
			{"Select a custom mode to delete", "freelance"},
			{"Delete \"freelance\" from saved accounts?", "No, keep it"},
			{"Please choose an account to configure", cancelSelectLabel},
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("HOME", t.TempDir())
			accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
			accounts := &models.Accounts{Custom: map[string]models.Account{
				"freelance": {Username: "freelancer", Email: "freelance@example.com"},
			}}
			if err := accountStorage.SaveAccounts(accounts); err != nil {
				t.Fatal(err)
			}
			setup := NewSetupService(NewAccountService(accountStorage, &testGit{}, &testSSH{})).(*SetupService)
			setup.selectOption = scriptedSetupSelections(t, tc.steps)
			if err := setup.SetupAccounts(); err != nil {
				t.Fatal(err)
			}
			saved, err := accountStorage.GetAccounts()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(saved, accounts) {
				t.Fatalf("cancellation changed accounts: %+v", saved)
			}
		})
	}
}

func TestSetupDeleteWithNoCustomModesReturnsToMenu(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
	setup := NewSetupService(NewAccountService(accountStorage, &testGit{}, &testSSH{})).(*SetupService)
	setup.selectOption = scriptedSetupSelections(t, []setupSelection{
		{"Please choose an account to configure", deleteSelectLabel},
		{"Please choose an account to configure", cancelSelectLabel},
	})
	if err := setup.SetupAccounts(); err != nil {
		t.Fatal(err)
	}
	saved, err := accountStorage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if len(saved.Custom) != 0 || saved.Work.Username != "" {
		t.Fatalf("unexpected accounts after empty deletion: %+v", saved)
	}
}

func TestSetupDirectDeletionUsesSameConfirmation(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
	accounts := &models.Accounts{Custom: map[string]models.Account{
		"freelance":   {Username: "freelancer", Email: "freelance@example.com"},
		"open-source": {Username: "contributor", Email: "oss@example.com"},
	}}
	if err := accountStorage.SaveAccounts(accounts); err != nil {
		t.Fatal(err)
	}
	setup := NewSetupService(NewAccountService(accountStorage, &testGit{}, &testSSH{})).(*SetupService)
	setup.selectOption = scriptedSetupSelections(t, []setupSelection{
		{"Delete \"freelance\" from saved accounts?", "Yes, delete it"},
	})
	if err := setup.DeleteCustomMode("freelance"); err != nil {
		t.Fatal(err)
	}
	saved, err := accountStorage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := saved.Custom["freelance"]; exists {
		t.Fatal("direct deletion left the mode in saved accounts")
	}
	if !reflect.DeepEqual(saved.Custom["open-source"], accounts.Custom["open-source"]) {
		t.Fatal("direct deletion changed another custom mode")
	}
}

func TestSetupDirectDeletionRejectsNonCustomMode(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
	accounts := &models.Accounts{Work: models.Account{Username: "worker", Email: "work@example.com"}}
	if err := accountStorage.SaveAccounts(accounts); err != nil {
		t.Fatal(err)
	}
	setup := NewSetupService(NewAccountService(accountStorage, &testGit{}, &testSSH{})).(*SetupService)
	setup.selectOption = func(label string, items []string) (int, error) {
		t.Fatalf("unexpected confirmation for %q", label)
		return 0, nil
	}
	if err := setup.DeleteCustomMode("work"); err == nil {
		t.Fatal("built-in mode could be deleted")
	}
	saved, err := accountStorage.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(saved, accounts) {
		t.Fatalf("built-in mode changed: %+v", saved)
	}
}
