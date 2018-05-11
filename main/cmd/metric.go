package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	command := cobra.Command{
		Use:              "metric [operation]",
		TraverseChildren: true,
		Short:            "Metrics-related commands.",
		Long:             "Metrics-related commands.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	command.AddCommand(InitMetricList())
	command.AddCommand(InitMetricPlot())
	rootCmd.AddCommand(&command)
}
