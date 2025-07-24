package cmd

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/connectors/ssh"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
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

		if savedAccounts.Personal.Username == currGitAccount.Username &&
			savedAccounts.Personal.Email == currGitAccount.Email &&
			(currGitAccount.SigningKeyID == "" || savedAccounts.Personal.SigningKeyID == currGitAccount.SigningKeyID) &&
			(currGitAccount.SSHKeyPath == "" || savedAccounts.Personal.SSHKeyPath == currGitAccount.SSHKeyPath) {
			logger.ReadCurrentAccountData(currGitAccount, models.PersonalMode)
			return
		}

		if savedAccounts.School.Username == currGitAccount.Username &&
			savedAccounts.School.Email == currGitAccount.Email &&
			(currGitAccount.SigningKeyID == "" || savedAccounts.School.SigningKeyID == currGitAccount.SigningKeyID) &&
			(currGitAccount.SSHKeyPath == "" || savedAccounts.School.SSHKeyPath == currGitAccount.SSHKeyPath) {
			logger.ReadCurrentAccountData(currGitAccount, models.SchoolMode)
			return
		}

		if savedAccounts.Work.Username == currGitAccount.Username &&
			savedAccounts.Work.Email == currGitAccount.Email &&
			(currGitAccount.SigningKeyID == "" || savedAccounts.Work.SigningKeyID == currGitAccount.SigningKeyID) &&
			(currGitAccount.SSHKeyPath == "" || savedAccounts.Work.SSHKeyPath == currGitAccount.SSHKeyPath) {
			logger.ReadCurrentAccountData(currGitAccount, models.WorkMode)
			return
		}

		isAccountSaved, err := accountService.CheckSavedAccount(currGitAccount)
		if err != nil {
			logger.PrintErrorExecutingMode()
			return
		}

		if !isAccountSaved {
			logger.ReadUnsavedGitAccount(currGitAccount)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
