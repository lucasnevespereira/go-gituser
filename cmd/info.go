package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about the accounts",
	Long:  "Print information about the accounts that you have configured",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		sshConnector := ssh.NewSSHConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector, sshConnector)
		savedAccounts, err := accountService.GetSavedAccounts()
		if err != nil {
			fmt.Println(err)
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}
		logger.ReadAccountsData(savedAccounts)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
