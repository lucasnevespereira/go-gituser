package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/external/git"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/external/repository/database"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/services/account"
	"strings"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Print the current account account",
	Long:  "Print the current account account you are using",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := database.Open(database.AccountsDBPath)
		if err != nil {
			logger.PrintError(err)
		}
		accountRepo := repository.NewAccountRepository(db)
		savedAccounts, err := accountRepo.GetAccounts()
		if err != nil {
			logger.PrintError(err)

		}

		gitService := git.NewService()
		accountService := account.NewAccountService(gitService)

		currEmail, currName := accountService.CurrentAccount()
		currName = strings.TrimSuffix(currName, "\n")
		currEmail = strings.TrimSuffix(currEmail, "\n")

		if savedAccounts.PersonalUsername == (currName) && savedAccounts.PersonalEmail == (currEmail) {
			logger.ReadCurrentAccountData(currName, currEmail, "personal")
			return
		}

		if savedAccounts.SchoolUsername == (currName) && savedAccounts.SchoolEmail == (currEmail) {
			logger.ReadCurrentAccountData(currName, currEmail, "school")
			return
		}

		if savedAccounts.WorkUsername == (currName) && savedAccounts.WorkEmail == (currEmail) {
			logger.ReadCurrentAccountData(currName, currEmail, "work")
			return
		}

		if accountService.UsernameIsUnsaved(savedAccounts, currName) || accountService.EmailIsUnsaved(savedAccounts, currEmail) {
			logger.ReadUnsavedGitAccount(currName, currEmail)
			return
		}

		if currName == "" || currEmail == "" {
			logger.PrintNoActiveMode()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
