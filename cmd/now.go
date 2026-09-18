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

		modes := matchingModes(savedAccounts, currGitAccount)
		switch len(modes) {
		case 0:
			logger.ReadUnsavedGitAccount(currGitAccount)
		case 1:
			logger.ReadCurrentAccountData(currGitAccount, modes[0])
		default:
			logger.ReadMatchingAccountsData(currGitAccount, modes)
		}
	},
}

func matchingModes(accounts *models.Accounts, current *models.Account) []string {
	var modes []string
	accounts.ForEachConfigured(func(mode string, account models.Account) bool {
		if account.Username == current.Username &&
			account.Email == current.Email &&
			(current.SigningKeyID == "" || account.SigningKeyID == current.SigningKeyID) &&
			(current.SSHKeyPath == "" || account.SSHKeyPath == current.SSHKeyPath) {
			modes = append(modes, mode)
		}
		return true
	})
	return modes
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
