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
	Short: "Manage SSH keys",
	Long:  "Manage SSH keys loaded in the SSH agent",
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
			fmt.Println("No SSH keys loaded in agent")
			return
		}

		fmt.Println("SSH keys in agent:")
		for _, key := range keys {
			fmt.Printf("  🔑 %s\n", key)
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
}
