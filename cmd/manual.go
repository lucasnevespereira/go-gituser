package cmd

import (
	"go-gituser/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var manualCmd = &cobra.Command{
	Use:   "manual",
	Short: "Print the manual",
	Long:  "Print a detailed manual about how to use the program",
	Run: func(cmd *cobra.Command, args []string) {
		logger.PrintManual()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(manualCmd)
}
