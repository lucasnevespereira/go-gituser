package cmd

import (
	"fmt"
	"go-gituser/internal/connectors/ssh"
	"go-gituser/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage SSH keys and configuration",
	Long: `Manage SSH keys and configuration for GitUser accounts.

This command provides tools to:
• List SSH keys currently loaded in the SSH agent
• Clear all SSH keys from the agent
• Discover existing SSH keys on your system
• Get help with SSH setup and configuration
• Test SSH connections to GitHub and GitLab`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show help when no subcommand is provided
		cmd.Help()
	},
}

var sshListCmd = &cobra.Command{
	Use:   "list",
	Short: "List SSH keys in agent",
	Long:  "List all SSH keys currently loaded in the SSH agent",
	Run: func(cmd *cobra.Command, args []string) {
		sshConnector := ssh.NewSSHConnector()
		keys, err := sshConnector.ListKeysInAgent()
		if err != nil {
			logger.PrintError(err)
			os.Exit(1)
		}

		if len(keys) == 0 {
			fmt.Println("ℹ️  No SSH keys loaded in agent")
			fmt.Println("💡 Tip: Switch to an account to load its SSH key: gituser work")
			return
		}

		fmt.Println("🔑 SSH keys currently loaded in agent:")
		for i, key := range keys {
			fmt.Printf("   %d. %s\n", i+1, key)
		}
	},
}

var sshClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all SSH keys from agent",
	Long:  "Remove all SSH keys from the SSH agent",
	Run: func(cmd *cobra.Command, args []string) {
		sshConnector := ssh.NewSSHConnector()
		if err := sshConnector.ClearAgent(); err != nil {
			logger.PrintError(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.AddCommand(sshListCmd)
	sshCmd.AddCommand(sshClearCmd)
	sshCmd.AddCommand(sshGuideCmd)
	sshCmd.AddCommand(sshDiscoverCmd)
	sshCmd.AddCommand(sshTestCmd)
}
