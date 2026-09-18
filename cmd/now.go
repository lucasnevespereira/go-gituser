package cmd

import (
	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/models"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Print the current git account",
	Long:  "Print the current git account you are using",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		sshConnector := ssh.NewSSHConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector, sshConnector)

		savedAccounts, err := accountService.GetSavedAccounts()
		if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}

		currGitAccount := accountService.GetCurrentGitAccount()
		if currGitAccount.Username == "" || currGitAccount.Email == "" {
			logger.PrintNoActiveMode()
			return
		}

		var activeMode string
		savedAccounts.ForEachConfigured(func(mode string, account models.Account) bool {
			if account.Username == currGitAccount.Username &&
				account.Email == currGitAccount.Email &&
				(currGitAccount.SigningKeyID == "" || account.SigningKeyID == currGitAccount.SigningKeyID) &&
				(currGitAccount.SSHKeyPath == "" || account.SSHKeyPath == currGitAccount.SSHKeyPath) {
				activeMode = mode
				return false
			}
			return true
		})
		if activeMode != "" {
			logger.ReadCurrentAccountData(currGitAccount, activeMode)
			return
		}
		logger.ReadUnsavedGitAccount(currGitAccount)
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
