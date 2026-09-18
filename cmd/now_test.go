package cmd

import (
	"reflect"
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
)

func TestMatchingModesReportsAmbiguousIdentities(t *testing.T) {
	accounts := &models.Accounts{Work: models.Account{
		Username: "shared", Email: "shared@example.com", SSHKeyPath: "/tmp/work-key",
	}}
	accounts.Set("freelance", models.Account{
		Username: "shared", Email: "shared@example.com", SSHKeyPath: "/tmp/freelance-key",
	})

	current := &models.Account{Username: "shared", Email: "shared@example.com"}
	if got, want := matchingModes(accounts, current), []string{"work", "freelance"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("matching modes = %v, want %v", got, want)
	}

	current.SSHKeyPath = "/tmp/freelance-key"
	if got, want := matchingModes(accounts, current), []string{"freelance"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("matching modes with loaded SSH key = %v, want %v", got, want)
	}
}
