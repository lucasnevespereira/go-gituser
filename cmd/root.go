package cmd

import (
	"github.com/spf13/cobra"
	"go-gituser/internal/pkg/logger"
	"go-gituser/internal/pkg/models"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "gituser",
	Short:   "Gituser - CLI to switch easily between account user account",
	Long:    "Gituser allows to switch between account accounts (work,student,personal)",
	Version: models.AppVersion,
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
