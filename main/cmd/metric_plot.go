package cmd

import (
	"github.com/spf13/cobra"
	"github.com/grafaneus"
	"encoding/json"
)

func InitMetricPlot() *cobra.Command {
	grafana := grafaneus.Grafana{}
	command := cobra.Command{
		Use:   "plot",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			metric, ok := grafaneus.MetricsMetadata[args[0]]
			var jsonx string
			if ok {
				jsonx = grafana.GenerateGraph(metric.Description, metric.Name)
			} else {
				jsonx = grafana.GenerateGraph(args[0], args[0])
			}
			var dash map[string]interface{}
			json.Unmarshal([]byte(jsonx), &dash)
			err := grafana.UploadDashboard(dash)
			if err != nil {
				panic(err)
			}
		},
	}
	return &command
}
