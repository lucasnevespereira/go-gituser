package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/services/git"
	"go-gituser/state"
	"go-gituser/utils"
	"go-gituser/utils/logger"
	"os"
)

var schoolCmd = &cobra.Command{
	Use:   "school",
	Short: "Switch to the school account",
	Long:  "Switch from your current account to the school account",
	Run: func(cmd *cobra.Command, args []string) {
		if state.SavedAccounts.SchoolUsername == "" {
			logger.PrintWarningReadingAccount(utils.SchoolMode)
			os.Exit(1)
		}
		git.SetAccount(state.SavedAccounts.SchoolUsername, state.SavedAccounts.SchoolEmail)
	},
}

func init() {
	rootCmd.AddCommand(schoolCmd)
}
