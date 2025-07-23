package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintManual() {
	fmt.Println("Hi there 👋🏼")
	fmt.Println("Welcome to GitUser - Your Git Account Manager 📘")
	fmt.Println("")

	color.Cyan("Description:")
	fmt.Println(" GitUser helps you switch between different git accounts seamlessly.")
	fmt.Println(" It manages Git configuration, GPG signing keys, and SSH keys automatically.")
	fmt.Println("")
	fmt.Println(" There are 3 account modes:")
	fmt.Println(" 🏠 | personal  - for your personal projects")
	fmt.Println(" 📚 | school    - for school/university work")
	fmt.Println(" 💻 | work      - for professional projects")
	fmt.Println("")

	color.Cyan("Features:")
	fmt.Println(" ✅ Switch Git username and email")
	fmt.Println(" ✅ Manage GPG signing keys")
	fmt.Println(" ✅ Handle SSH keys automatically")
	fmt.Println(" ✅ Seamless GitHub/GitLab authentication")
	fmt.Println("")

	color.Cyan("Quick Start:")
	fmt.Println(" 1. Run initial setup:")
	fmt.Println("    gituser setup")
	fmt.Println("")
	fmt.Println(" 2. Switch to an account:")
	fmt.Println("    gituser work")
	fmt.Println("    gituser personal")
	fmt.Println("    gituser school")
	fmt.Println("")

	color.Cyan("Setup Process:")
	fmt.Println(" The setup wizard will guide you through:")
	fmt.Println(" • Adding your username and email for each account")
	fmt.Println(" • Configuring GPG signing keys (optional)")
	fmt.Println(" • Setting up SSH keys with these options:")
	fmt.Println("   - Auto-detect existing SSH keys")
	fmt.Println("   - Generate new SSH keys")
	fmt.Println("   - Manually specify key paths")
	fmt.Println("   - Skip SSH setup if not needed")
	fmt.Println("")

	color.Cyan("Account Management:")
	fmt.Println(" gituser setup       - Configure your accounts")
	fmt.Println(" gituser info        - View all configured accounts")
	fmt.Println(" gituser now         - Show currently active account")
	fmt.Println("")

	color.Cyan("SSH Key Management:")
	fmt.Println(" gituser ssh list      - List SSH keys in agent")
	fmt.Println(" gituser ssh clear     - Clear all SSH keys from agent")
	fmt.Println(" gituser ssh discover  - Find existing SSH keys")
	fmt.Println(" gituser ssh guide     - Show SSH setup guide")
	fmt.Println(" gituser ssh test      - Test GitHub/GitLab connections")
	fmt.Println("")

	color.Cyan("Usage Examples:")
	fmt.Println(" # Switch to work account for company projects")
	fmt.Println(" gituser work")
	fmt.Println("")
	fmt.Println(" # Switch to personal account for side projects")
	fmt.Println(" gituser personal")
	fmt.Println("")
	fmt.Println(" # Check which account is currently active")
	fmt.Println(" gituser now")
	fmt.Println("")
	fmt.Println(" # View all your configured accounts")
	fmt.Println(" gituser info")
	fmt.Println("")

	color.Cyan("Help & Support:")
	fmt.Println(" gituser help          - Show available commands")
	fmt.Println(" gituser ssh help      - SSH-specific help")
	fmt.Println(" gituser ssh guide     - Detailed SSH setup guide")
	fmt.Println("")

	color.Yellow("💡 Pro Tips:")
	fmt.Println(" • Use descriptive SSH key names (e.g., id_ed25519_work)")
	fmt.Println(" • Test your SSH connections after setup")
	fmt.Println(" • Run 'gituser now' to verify you're on the right account")
	fmt.Println(" • SSH keys are automatically managed when switching accounts")
	fmt.Println("")
}
