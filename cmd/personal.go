package cmd

import (
	"errors"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var personalCmd = &cobra.Command{
	Use:   "personal",
	Short: "Switch to the personal account",
	Long:  "Switch from your current account to the personal account",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		sshConnector := ssh.NewSSHConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector, sshConnector)
		err := accountService.Switch(models.PersonalMode)
		if err != nil && errors.Is(err, models.ErrNoAccountFound) {
			logger.PrintWarningReadingAccount(models.PersonalMode)
			os.Exit(1)
		} else if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(personalCmd)
}
