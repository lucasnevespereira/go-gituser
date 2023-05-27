package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/external/git"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/external/repository/database"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/services/account"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Switch to the work account",
	Long:  "Switch from your current account to the work account",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := database.Open(database.AccountsDBPath)
		if err != nil {
			logger.PrintError(err)
		}
		accountRepo := repository.NewAccountRepository(db)

		accounts, err := accountRepo.GetAccounts()
		if err != nil {
			logger.PrintError(err)
		}

		gitService := git.NewService()
		accountService := account.NewAccountService(gitService)
		accountService.SetAccount(accounts.WorkUsername, accounts.WorkEmail)
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
