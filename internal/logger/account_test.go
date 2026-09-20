package logger

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
)

func captureAccountsOutput(t *testing.T, accounts *models.Accounts) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	previousStdout := os.Stdout
	os.Stdout = writer
	defer func() {
		os.Stdout = previousStdout
		reader.Close()
		writer.Close()
	}()

	ReadAccountsData(accounts)
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	return string(output)
}

func TestReadAccountsDataShowsOnlyConfiguredAccountsInStableOrder(t *testing.T) {
	accounts := &models.Accounts{
		Personal: models.Account{Username: "alice", Email: "alice@example.com", SigningKeyID: "personal-gpg", SSHKeyPath: "/keys/personal"},
		Work:     models.Account{Username: "bob", Email: "bob@example.com"},
		Custom: map[string]models.Account{
			"zeta":      {Username: "zoe", Email: "zoe@example.com", SSHKeyPath: "/keys/zeta"},
			"freelance": {Username: "carol", Email: "carol@example.com", SigningKeyID: "freelance-gpg"},
			"unused":    {},
		},
	}
	output := captureAccountsOutput(t, accounts)

	headings := []string{
		"🏠 | Personal Git Account :",
		"💻 | Work Git Account :",
		"🔖 | freelance Git Account :",
		"🔖 | zeta Git Account :",
	}
	previous := -1
	for _, heading := range headings {
		position := strings.Index(output, heading)
		if position <= previous {
			t.Fatalf("account %q missing or out of order in output:\n%s", heading, output)
		}
		previous = position
	}
	for _, field := range []string{"alice@example.com", "bob@example.com", "carol@example.com", "zoe@example.com", "personal-gpg", "/keys/personal", "freelance-gpg", "/keys/zeta"} {
		if !strings.Contains(output, field) {
			t.Errorf("configured field %q missing from output:\n%s", field, output)
		}
	}
	for _, unwanted := range []string{"You have no", "School Git Account", "🔖 | unused Git Account"} {
		if strings.Contains(output, unwanted) {
			t.Errorf("unconfigured account %q appeared in output:\n%s", unwanted, output)
		}
	}
}

func TestReadAccountsDataShowsOneEmptyStateMessage(t *testing.T) {
	accounts := &models.Accounts{Custom: map[string]models.Account{"unused": {}}}
	output := captureAccountsOutput(t, accounts)
	want := "No accounts configured. Run gituser setup to add one.\n"
	if output != want {
		t.Fatalf("empty account output = %q, want %q", output, want)
	}
}
