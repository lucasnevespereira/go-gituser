package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/external/git"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/external/repository/database"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/services/account"
)

var personalCmd = &cobra.Command{
	Use:   "personal",
	Short: "Switch to the personal account",
	Long:  "Switch from your current account to the personal account",
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
		accountService.SetAccount(accounts.PersonalUsername, accounts.PersonalEmail)
	},
}

func init() {
	rootCmd.AddCommand(personalCmd)
}
