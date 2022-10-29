package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/app"
	"go-gituser/internal/models"
	"go-gituser/state"
	"go-gituser/utils"
	"os"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about the accounts",
	Long:  "Print information about the accounts that you have configured",
	Run: func(cmd *cobra.Command, args []string) {

		app.Sync()

		var currentAccounts models.Accounts
		if state.SavedAccounts != nil {
			currentAccounts = *state.SavedAccounts
		}

		utils.ReadAccountsData(currentAccounts)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
