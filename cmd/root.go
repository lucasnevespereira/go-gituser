package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/utils"
	"go-gituser/utils/logger"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gituser",
	Short: "Gituser - CLI to switch easily between git user account",
	Long:  "Gituser allows to switch between git accounts (work,student,personal)",
	Version: utils.AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
