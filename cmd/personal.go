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

var personalCmd = &cobra.Command{
	Use:   "personal",
	Short: "Switch to the personal account",
	Long:  "Switch from your current account to the personal account",
	Run: func(cmd *cobra.Command, args []string) {
		app.Sync()
		if state.SavedAccounts.PersonalUsername == "" {
			logger.PrintWarningReadingAccount(utils.PersonalMode)
			os.Exit(1)
		}
		git.SetAccount(state.SavedAccounts.PersonalUsername, state.SavedAccounts.PersonalEmail)
	},
}

func init() {
	rootCmd.AddCommand(personalCmd)
}
