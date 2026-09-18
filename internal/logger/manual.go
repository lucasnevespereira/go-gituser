package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintManual() {
	fmt.Println("🔄 GitUser - Switch Between Git Accounts Easily")
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()

	color.Cyan("🚀 Quick Start (30 seconds):")
	fmt.Println("  1. gituser setup     # Configure your accounts")
	fmt.Println("  2. gituser work      # Switch to work account")
	fmt.Println("  3. gituser personal  # Switch to personal account")
	fmt.Println("  4. gituser now       # Check current account")
	fmt.Println()

	color.Cyan("✨ What GitUser Does:")
	fmt.Println("  Instead of manually running:")
	fmt.Println("    git config --global user.name \"Work Name\"")
	fmt.Println("    git config --global user.email \"work@company.com\"")
	fmt.Println("    ssh-add ~/.ssh/work_key")
	fmt.Println()
	fmt.Println("  Just run: gituser work ✨")
	fmt.Println()

	color.Cyan("🏠 Your Account Modes:")
	fmt.Println("  💻 work     - for company projects")
	fmt.Println("  🏠 personal - for side projects")
	fmt.Println("  📚 school   - for university work")
	fmt.Println("  🔖 custom   - create your own modes in gituser setup")
	fmt.Println()

	color.Cyan("📋 Essential Commands:")
	fmt.Println("  gituser setup       # Configure accounts (run once)")
	fmt.Println("  gituser work        # Switch to work")
	fmt.Println("  gituser personal    # Switch to personal")
	fmt.Println("  gituser school      # Switch to school")
	fmt.Println("  gituser freelance   # Switch to a configured custom mode")
	fmt.Println("  gituser now         # Show current account")
	fmt.Println("  gituser info        # View all accounts")
	fmt.Println()

	color.Cyan("🔑 SSH & Security:")
	fmt.Println("  gituser ssh discover # Find your SSH keys")
	fmt.Println("  gituser ssh test     # Test GitHub/GitLab")
	fmt.Println("  gituser ssh guide    # SSH setup help")
	fmt.Println()

	color.Yellow("💡 Pro Tips:")
	fmt.Println("  • Always run 'gituser now' to check your current account")
	fmt.Println("  • SSH keys automatically switch with accounts")
	fmt.Println("  • Setup wizard handles everything - GPG, SSH, Git config")
	fmt.Println()

	color.Green("🎯 Next Steps:")
	fmt.Println("  New user? → gituser setup")
	fmt.Println("  Need help? → gituser help")
	fmt.Println("  SSH issues? → gituser ssh guide")
	fmt.Println()
}
