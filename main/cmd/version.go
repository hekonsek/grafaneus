package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Grafaneus version",
	Long:  `Print the version number of Grafaneus`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0")
	},
}