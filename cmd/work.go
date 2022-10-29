package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/app"
	"go-gituser/internal/services/git"
	"go-gituser/state"
	"go-gituser/utils"
	"go-gituser/utils/logger"
	"os"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Switch to the work account",
	Long:  "Switch from your current account to the work account",
	Run: func(cmd *cobra.Command, args []string) {
		app.Sync()

		if state.SavedAccounts.WorkUsername == "" {
			logger.PrintWarningReadingAccount(utils.WorkMode)
			os.Exit(1)
		}
		git.SetAccount(state.SavedAccounts.WorkUsername, state.SavedAccounts.WorkEmail)
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
