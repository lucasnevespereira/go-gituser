package cmd

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/connectors/ssh"
	"go-gituser/internal/logger"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your different git accounts",
	Long:  "Modify or init the configuration (email,username) of your different git accounts (work,school,personal)",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		sshConnector := ssh.NewSSHConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector, sshConnector)
		setupService := services.NewSetupService(accountService)
		if err := setupService.SetupAccounts(); err != nil {
			logger.PrintError(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
