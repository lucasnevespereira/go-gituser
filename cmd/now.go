package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/app"
	"go-gituser/internal/services/git"
	"go-gituser/state"
	"go-gituser/utils"
	"go-gituser/utils/logger"
	"strings"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Print the current git account",
	Long:  "Print the current git account you are using",
	Run: func(cmd *cobra.Command, args []string) {
		app.Sync()

		currEmail, currName := git.CurrentAccount()
		currName = strings.TrimSuffix(currName, "\n")
		currEmail = strings.TrimSuffix(currEmail, "\n")

		if state.SavedAccounts.PersonalUsername == (currName) && state.SavedAccounts.PersonalEmail == (currEmail) {
			utils.ReadCurrentAccountData(currName, currEmail, "personal")
			return
		}

		if state.SavedAccounts.SchoolUsername == (currName) && state.SavedAccounts.SchoolEmail == (currEmail) {
			utils.ReadCurrentAccountData(currName, currEmail, "school")
			return
		}

		if state.SavedAccounts.WorkUsername == (currName) && state.SavedAccounts.WorkEmail == (currEmail) {
			utils.ReadCurrentAccountData(currName, currEmail, "work")
			return
		}

		if utils.GitUsernameIsUnsaved(currName) || utils.GitEmailIsUnsaved(currEmail) {
			utils.ReadUnsavedGitAccount(currName, currEmail)
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
