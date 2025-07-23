package cmd

import (
	"go-gituser/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gituser",
	Short: "GitUser - Switch between git accounts easily",
	Long: `🔄 GitUser - Your Complete Git Account Manager

GitUser helps you seamlessly switch between multiple Git accounts by managing:
• Git username and email configuration
• GPG signing keys for verified commits
• SSH keys for secure authentication
• Account switching with a single command

Perfect for developers who work with multiple accounts (work, personal, school)
and need to switch between them frequently without manual configuration.

Quick Start:
  gituser setup          # Interactive setup wizard
  gituser work           # Switch to work account
  gituser personal       # Switch to personal account
  gituser now            # Show current active account
  gituser ssh discover   # Find your SSH keys

Get Help:
  gituser help           # Show all commands
  gituser ssh help       # SSH-specific commands
  gituser manual         # Detailed manual`,
	Version: AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
