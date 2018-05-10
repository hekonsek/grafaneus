package cmd

import (
	"github.com/spf13/cobra"
	"github.com/grafaneus"
	"fmt"
)

func InitMetricPlot() *cobra.Command {
	grafana := grafaneus.Grafana{}
	command := cobra.Command{
		Use:   "plot",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			metric, ok := grafaneus.MetricsMetadata[args[0]]
			if ok {
				fmt.Println(grafana.GenerateGraph(metric.Description, metric.Name))
			} else {
				fmt.Println(grafana.GenerateGraph(args[0], args[0]))
			}
		},
	}
	return &command

}