package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/external/repository"
	"go-gituser/internal/pkg/external/repository/database"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/services/setup"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your different account accounts",
	Long:  "Modify or init the configuration (email,username) of your different account accounts (work,school,personal)",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.Open(database.AccountsDBPath)
		if err != nil {
			logger.PrintError(err)
		}
		accountRepo := repository.NewAccountRepository(db)
		setupService := setup.NewSetupService(accountRepo)

		setupService.SetupAccounts()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
