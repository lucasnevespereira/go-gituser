package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/app"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup your different git accounts",
	Long:  "Modify or init the configuration (email,username) of your different git accounts (work,school,personal)",
	Run: func(cmd *cobra.Command, args []string) {
		app.SetupAccounts()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
