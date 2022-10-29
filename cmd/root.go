package cmd

import (
	"github.com/spf13/cobra"
	logger2 "go-gituser/utils/logger"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gituser",
	Short: "Gituser - a CLI to switch easily between git user account",
	Long:  "Gituser allows to switch between git accounts (work,student,personal)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logger2.PrintManual()
			os.Exit(1)
		}

		if len(args) > 2 {
			logger2.PrintErrorInvalidArguments()
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger2.PrintError(err)
		os.Exit(1)
	}
}
