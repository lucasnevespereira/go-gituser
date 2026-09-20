package cmd

import (
	"github.com/lucasnevespereira/go-gituser/internal/connectors/git"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your different git accounts",
	Long:  "Configure built-in or custom git accounts (email, username, GPG, SSH)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := newSetupService().SetupAccounts(); err != nil {
			logger.PrintError(err)
			os.Exit(1)
		}
	},
}

var setupRemoveCmd = &cobra.Command{
	Use:   "remove <mode>",
	Short: "Remove a saved custom mode",
	Long:  "Remove a saved custom mode without deleting its SSH key or changing the current Git configuration.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return newSetupService().DeleteCustomMode(args[0])
	},
}

func newSetupService() services.ISetupService {
	accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
	accountService := services.NewAccountService(accountStorage, git.NewGitConnector(), ssh.NewSSHConnector())
	return services.NewSetupService(accountService)
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.AddCommand(setupRemoveCmd)
}
