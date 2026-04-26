package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/go-gituser/internal/connectors/ssh"
	"github.com/lucasnevespereira/go-gituser/internal/services"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var sshGuideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Show SSH setup guide",
	Long:  "Display a comprehensive guide for SSH key setup",
	Run: func(cmd *cobra.Command, args []string) {
		sshConnector := ssh.NewSSHConnector()
		sshDiscovery := services.NewSSHDiscoveryService(sshConnector)
		sshDiscovery.ShowSSHSetupGuide()
	},
}

var sshDiscoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover existing SSH keys",
	Long:  "Scan for and display existing SSH keys on your system",
	Run: func(cmd *cobra.Command, args []string) {
		sshConnector := ssh.NewSSHConnector()
		sshDiscovery := services.NewSSHDiscoveryService(sshConnector)

		keys, err := sshDiscovery.DiscoverSSHKeys()
		if err != nil {
			fmt.Printf("❌ Error discovering SSH keys: %v\n", err)
			return
		}

		fmt.Println("🔍 SSH Key Discovery Results")
		fmt.Println("═══════════════════════════════════════════════════════════════")
		fmt.Println()

		existingKeys := make([]services.SSHKeyInfo, 0)
		missingKeys := make([]services.SSHKeyInfo, 0)

		for _, key := range keys {
			if key.Exists {
				existingKeys = append(existingKeys, key)
			} else {
				missingKeys = append(missingKeys, key)
			}
		}

		if len(existingKeys) > 0 {
			fmt.Println("✅ FOUND SSH KEYS")
			fmt.Println("─────────────────")
			fmt.Println()

			for i, key := range existingKeys {
				// Determine status and color
				var statusIcon, statusText, keyIcon string
				if key.HasPublic {
					statusIcon = "✅"
					statusText = "READY TO USE"
					keyIcon = "🔑"
				} else {
					statusIcon = "⚠️ "
					statusText = "MISSING PUBLIC KEY"
					keyIcon = "🔓"
				}

				fmt.Printf("%s %s SSH Key #%d\n", keyIcon, key.Type, i+1)
				fmt.Printf("   Status:      %s %s\n", statusIcon, statusText)
				fmt.Printf("   Private Key: %s\n", key.Path)

				if key.HasPublic {
					fmt.Printf("   Public Key:  %s.pub\n", key.Path)
				} else {
					fmt.Printf("   Public Key:  %s.pub (❌ NOT FOUND)\n", key.Path)
				}

				fmt.Printf("   Description: %s\n", key.Description)

				if !key.CreatedTime.IsZero() {
					fmt.Printf("   Created:     %s\n", key.CreatedTime.Format("January 2, 2006 at 3:04 PM"))
				}

				// Add usage hint for keys missing public keys
				if !key.HasPublic {
					fmt.Printf("   💡 Generate public key: ssh-keygen -y -f %s > %s.pub\n", key.Path, key.Path)
				}

				fmt.Println()
			}

			fmt.Println("Legend:")
			fmt.Println("  ✅ Ready to use (has both private and public key)")
			fmt.Println("  ⚠️  Missing public key (needs to be generated)")
			fmt.Println()
		}

		if len(missingKeys) > 0 {
			fmt.Println("💡 COMMON SSH KEY LOCATIONS (not found)")
			fmt.Println("──────────────────────────────────────────")
			fmt.Println()

			for _, key := range missingKeys {
				var recommendation string
				switch key.Type {
				case "ED25519":
					recommendation = "🌟 RECOMMENDED - Modern, secure, and fast"
				case "RSA":
					recommendation = "👍 COMPATIBLE - Widely supported"
				case "ECDSA":
					recommendation = "⚡ EFFICIENT - Good balance of security and size"
				case "DSA":
					recommendation = "❌ DEPRECATED - Not recommended for new keys"
				}

				fmt.Printf("📁 %s\n", key.Path)
				fmt.Printf("   Type:        %s\n", key.Type)
				fmt.Printf("   Status:      %s\n", recommendation)
				fmt.Printf("   Description: %s\n", key.Description)
				fmt.Println()
			}
		}

		if len(existingKeys) == 0 && len(missingKeys) > 0 {
			fmt.Println("🆕 NO EXISTING SSH KEYS FOUND")
			fmt.Println("─────────────────────────────")
			fmt.Println()
			fmt.Println("To get started with SSH keys:")
			fmt.Println("  1. Run: gituser setup")
			fmt.Println("  2. Choose to generate new SSH keys when prompted")
			fmt.Println("  3. Or create them manually:")
			fmt.Println("     • Ed25519 (recommended): ssh-keygen -t ed25519 -C \"your@email.com\"")
			fmt.Println("     • RSA (compatible):      ssh-keygen -t rsa -b 4096 -C \"your@email.com\"")
			fmt.Println()
		}

		// Summary
		if len(existingKeys) > 0 {
			readyCount := 0
			for _, key := range existingKeys {
				if key.HasPublic {
					readyCount++
				}
			}

			fmt.Println("📊 SUMMARY")
			fmt.Println("──────────")
			fmt.Printf("  Total SSH keys found: %d\n", len(existingKeys))
			fmt.Printf("  Ready to use:         %d\n", readyCount)
			fmt.Printf("  Need public key:      %d\n", len(existingKeys)-readyCount)
			fmt.Println()

			if readyCount > 0 {
				fmt.Println("🎉 You have SSH keys ready! Run 'gituser setup' to configure them with your accounts.")
			} else {
				fmt.Println("⚠️  Your SSH keys need public keys generated before they can be used.")
			}
		}
	},
}

var sshTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test SSH connections",
	Long:  "Test SSH connections to GitHub and GitLab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧪 Testing SSH Connections")
		fmt.Println("===========================")

		services := []struct {
			name string
			host string
		}{
			{"GitHub", "git@github.com"},
			{"GitLab", "git@gitlab.com"},
		}

		for _, service := range services {
			fmt.Printf("\n🔗 Testing %s connection...\n", service.name)

			cmd := exec.Command("ssh", "-T", "-o", "StrictHostKeyChecking=no", service.host)
			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			if strings.Contains(outputStr, "successfully authenticated") ||
				strings.Contains(outputStr, "You've successfully authenticated") {
				fmt.Printf("   ✅ %s connection successful!\n", service.name)
			} else if err != nil {
				fmt.Printf("   ❌ %s connection failed\n", service.name)
				fmt.Printf("   Error: %s\n", outputStr)
			} else {
				fmt.Printf("   ⚠️  %s connection unclear: %s\n", service.name, outputStr)
			}
		}

		fmt.Println("\n💡 If connections failed:")
		fmt.Println("   1. Make sure you've added your public key to the service")
		fmt.Println("   2. Check that your SSH key is loaded: gituser ssh list")
		fmt.Println("   3. Run setup again: gituser setup")
	},
}
