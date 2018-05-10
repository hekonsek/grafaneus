package cmd

import (
	"github.com/spf13/cobra"
	"github.com/grafaneus"
	"fmt"
)

func init() {
	command := cobra.Command{
		Use:   "metric-plot",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			metric, ok := grafaneus.MetricsMetadata[args[0]]
			if ok {
				fmt.Println(grafaneus.GenerateGraph(metric.Description, metric.Name))
			} else {
				fmt.Println(grafaneus.GenerateGraph(args[0], args[0]))
			}
		},
	}
	rootCmd.AddCommand(&command)
}