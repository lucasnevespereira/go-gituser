package cmd

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/logger"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about the accounts",
	Long:  "Print information about the accounts that you have configured",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector)
		savedAccounts, err := accountService.ReadSavedAccounts()
		if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}
		logger.ReadAccountsData(savedAccounts)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
