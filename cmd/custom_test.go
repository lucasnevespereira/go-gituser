package cmd

import (
	"testing"

	"github.com/lucasnevespereira/go-gituser/internal/models"
)

func TestRegisterCustomModes(t *testing.T) {
	accounts := &models.Accounts{Custom: map[string]models.Account{
		"freelance": {Username: "freelancer", Email: "freelance@example.com"},
		"remove":    {Username: "remover", Email: "remove@example.com"},
		"setup":     {Username: "invalid", Email: "invalid@example.com"},
		"empty":     {},
	}}
	registerCustomModes(accounts)

	command, _, err := rootCmd.Find([]string{"freelance"})
	if err != nil || command.Name() != "freelance" {
		t.Fatalf("custom command = %v, %v", command, err)
	}
	command, _, err = rootCmd.Find([]string{"setup"})
	if err != nil || command != setupCmd {
		t.Fatalf("setup command was replaced: %v, %v", command, err)
	}
	command, _, err = rootCmd.Find([]string{"remove"})
	if err != nil || command.Name() != "remove" {
		t.Fatalf("custom remove mode = %v, %v", command, err)
	}
	command, _, err = rootCmd.Find([]string{"setup", "remove"})
	if err != nil || command != setupRemoveCmd {
		t.Fatalf("setup remove command = %v, %v", command, err)
	}
	for _, command := range rootCmd.Commands() {
		if command.Name() == "empty" {
			t.Fatal("empty custom mode was registered")
		}
	}
}
