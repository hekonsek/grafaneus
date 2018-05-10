package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	command := cobra.Command{
		Use:   "metric [operation]",
		TraverseChildren:true,
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Metric root")
		},
	}
	command.AddCommand(InitMetricList())
	command.AddCommand(InitMetricPlot())
	rootCmd.AddCommand(&command)
}