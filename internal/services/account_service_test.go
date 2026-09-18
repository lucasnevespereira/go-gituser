package services

import (
	"errors"
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
)

type testStorage struct{ accounts *models.Accounts }

func (s *testStorage) GetAccounts() (*models.Accounts, error) { return s.accounts, nil }
func (s *testStorage) GetAccountByUsername(username string) (*models.Account, error) {
	return nil, errors.New("not found")
}
func (s *testStorage) SaveAccounts(accounts *models.Accounts) error {
	s.accounts = accounts
	return nil
}

type testGit struct {
	configured *models.Account
	current    models.Account
}

func (g *testGit) ReadConfig() *models.Account {
	copy := g.current
	return &copy
}
func (g *testGit) SetConfig(account *models.Account) {
	copy := *account
	g.configured = &copy
}

type testSSH struct {
	cleared int
	added   string
	loaded  map[string]bool
}

func (s *testSSH) AddKeyToAgent(path string) error            { s.added = path; return nil }
func (s *testSSH) RemoveKeyFromAgent(string) error            { return nil }
func (s *testSSH) ListKeysInAgent() ([]string, error)         { return nil, nil }
func (s *testSSH) ClearAgent() error                          { s.cleared++; return nil }
func (s *testSSH) IsKeyLoaded(path string) bool               { return s.loaded[path] }
func (s *testSSH) ValidateKeyPath(string) error               { return nil }
func (s *testSSH) GetDefaultKeyPath() string                  { return "" }
func (s *testSSH) StartSSHAgent() error                       { return nil }
func (s *testSSH) GetPublicKeyContent(string) (string, error) { return "", nil }

func TestSwitchCustomMode(t *testing.T) {
	accounts := &models.Accounts{Work: models.Account{Username: "worker", Email: "work@example.com"}}
	accounts.Set("freelance", models.Account{Username: "freelancer", Email: "freelance@example.com", SSHKeyPath: "/tmp/freelance-key"})
	git := &testGit{}
	ssh := &testSSH{}
	service := NewAccountService(&testStorage{accounts}, git, ssh)

	if err := service.Switch("missing"); !errors.Is(err, models.ErrNoAccountFound) {
		t.Fatalf("missing mode error = %v", err)
	}
	if git.configured != nil || ssh.cleared != 0 {
		t.Fatal("missing mode changed Git config or SSH agent")
	}
	if err := service.Switch("freelance"); err != nil {
		t.Fatal(err)
	}
	if git.configured == nil || git.configured.Username != "freelancer" || ssh.cleared != 1 || ssh.added != "/tmp/freelance-key" {
		t.Fatalf("switch did not use custom account: Git=%+v SSH=%+v", git.configured, ssh)
	}
	if saved, err := service.CheckSavedAccount(&models.Account{Username: "freelancer", Email: "freelance@example.com"}); err != nil || !saved {
		t.Fatalf("custom account not recognized as saved: saved=%v err=%v", saved, err)
	}
}

func TestCurrentGitAccountUsesMatchingIdentityAndLoadedSSHKey(t *testing.T) {
	accounts := &models.Accounts{Work: models.Account{
		Username: "shared", Email: "work@example.com", SSHKeyPath: "/tmp/work-key",
	}}
	accounts.Set("freelance", models.Account{
		Username: "shared", Email: "freelance@example.com", SSHKeyPath: "/tmp/freelance-key",
	})
	git := &testGit{current: models.Account{Username: "shared", Email: "freelance@example.com"}}
	ssh := &testSSH{loaded: map[string]bool{"/tmp/freelance-key.pub": true}}
	service := NewAccountService(&testStorage{accounts}, git, ssh)

	current := service.GetCurrentGitAccount()
	if current.SSHKeyPath != "/tmp/freelance-key" {
		t.Fatalf("current SSH key = %q, want freelance key", current.SSHKeyPath)
	}
}
