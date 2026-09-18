package cmd

import (
	"github.com/lucasnevespereira/go-gituser/internal/logger"
	"github.com/lucasnevespereira/go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gituser",
	Short: "GitUser - Switch between git accounts easily",
	Long: `GitUser - Switch between your Git accounts with one command

Perfect for developers with multiple Git accounts (work, personal, school).
Automatically manages Git config, GPG keys, and SSH keys.

Quick Start:
  gituser setup     # Configure your accounts
  gituser work      # Switch to work account
  gituser personal  # Switch to personal account
  gituser school    # Switch to student account
  gituser freelance # Switch to a custom account after setup
  gituser now       # Check current account

Need help? Run: gituser help`,
	Version: AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
	},
}

func Execute() {
	accounts, err := storage.NewAccountJSONStorage(storage.AccountsStorageFile).GetAccounts()
	if err == nil {
		registerCustomModes(accounts)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
