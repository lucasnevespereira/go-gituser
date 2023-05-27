package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/external/repository/database"
	"go-gituser/internal/pkg/logger"
	"os"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about the accounts",
	Long:  "Print information about the accounts that you have configured",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := database.Open(database.AccountsDBPath)
		if err != nil {
			logger.PrintError(err)
		}
		accountRepo := repository.NewAccountRepository(db)

		currentAccounts, err := accountRepo.GetAccounts()
		if err != nil {
			logger.PrintError(err)

		}

		logger.ReadAccountsData(currentAccounts)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
