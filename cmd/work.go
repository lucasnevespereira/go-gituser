package cmd

import (
	"errors"
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/connectors/ssh"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Switch to the work account",
	Long:  "Switch from your current account to the work account",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		sshConnector := ssh.NewSSHConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector, sshConnector)
		err := accountService.Switch(models.WorkMode)
		if err != nil && errors.Is(err, models.ErrNoAccountFound) {
			logger.PrintWarningReadingAccount(models.WorkMode)
			os.Exit(1)
		} else if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
