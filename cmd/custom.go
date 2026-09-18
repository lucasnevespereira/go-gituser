package cmd

import (
	"errors"
	"os"

	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"github.com/spf13/cobra"
)

func registerCustomModes(accounts *models.Accounts) {
	for mode, account := range accounts.Custom {
		if account.Username == "" || models.ValidateCustomMode(mode) != nil {
			continue
		}
		reserved := false
		for _, command := range rootCmd.Commands() {
			if command.Name() == mode {
				reserved = true
				break
			}
		}
		if reserved {
			continue
		}

		mode := mode
		rootCmd.AddCommand(&cobra.Command{
			Use:   mode,
			Short: "Switch to the " + mode + " account",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
				accountService := services.NewAccountService(accountStorage, git.NewGitConnector(), ssh.NewSSHConnector())
				if err := accountService.Switch(mode); err != nil {
					if errors.Is(err, models.ErrNoAccountFound) {
						logger.PrintWarningReadingAccount(mode)
					} else {
						logger.PrintErrorExecutingMode()
					}
					os.Exit(1)
				}
			},
		})
	}
}
