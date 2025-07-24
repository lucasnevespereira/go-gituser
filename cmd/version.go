package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AppVersion = "0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print gituser version",
	Long:  "Print version number of gituser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gituser version %s \n", AppVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
