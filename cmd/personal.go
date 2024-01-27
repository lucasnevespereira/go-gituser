package cmd

import (
	"errors"
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
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
		accountService := services.NewAccountService(accountStorage, gitConnector)
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
