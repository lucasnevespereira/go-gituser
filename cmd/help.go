package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/logger"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "manual",
	Short: "Show detailed manual",
	Long:  "Display the complete GitUser manual with examples and tips",
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
	},
}

var quickHelpCmd = &cobra.Command{
	Use:   "quickstart",
	Short: "Show quick start guide",
	Long:  "Display a quick start guide to get you up and running in 30 seconds",
	Run: func(cmd *cobra.Command, args []string) {
		printQuickStart()
	},
}

func printQuickStart() {
	fmt.Println("🚀 GitUser Quick Start")
	fmt.Println("═══════════════════════")
	fmt.Println()
	fmt.Println("1️⃣ Setup your accounts:")
	fmt.Println("   gituser setup")
	fmt.Println()
	fmt.Println("2️⃣ Switch accounts:")
	fmt.Println("   gituser work      # Switch to work")
	fmt.Println("   gituser personal  # Switch to personal")
	fmt.Println("   gituser school    # Switch to school")
	fmt.Println()
	fmt.Println("3️⃣ Check current account:")
	fmt.Println("   gituser now")
	fmt.Println()
	fmt.Println("That's it! 🎉")
	fmt.Println()
	fmt.Println("Need more help? Run: gituser manual")
}

func init() {
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(quickHelpCmd)
}
