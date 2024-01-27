package cmd

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Print the current git account",
	Long:  "Print the current git account you are using",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector)

		savedAccounts, err := accountService.ReadSavedAccounts()
		if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}

		currGitUsername, currGitEmail := accountService.ReadCurrentGitAccount()
		if currGitUsername == "" || currGitEmail == "" {
			logger.PrintNoActiveMode()
			return
		}

		if savedAccounts.PersonalUsername == (currGitUsername) && savedAccounts.PersonalEmail == (currGitEmail) {
			logger.ReadCurrentAccountData(currGitUsername, currGitEmail, models.PersonalMode)
			return
		}

		if savedAccounts.SchoolUsername == (currGitUsername) && savedAccounts.SchoolEmail == (currGitEmail) {
			logger.ReadCurrentAccountData(currGitUsername, currGitEmail, models.SchoolMode)
			return
		}

		if savedAccounts.WorkUsername == (currGitUsername) && savedAccounts.WorkEmail == (currGitEmail) {
			logger.ReadCurrentAccountData(currGitUsername, currGitEmail, models.WorkMode)
			return
		}

		isAccountSaved, err := accountService.CheckSavedAccount(currGitUsername, currGitEmail)
		if err != nil {
			logger.PrintErrorExecutingMode()
			return
		}

		if !isAccountSaved {
			logger.ReadUnsavedGitAccount(currGitUsername, currGitEmail)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
