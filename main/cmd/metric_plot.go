package cmd

import (
	"github.com/spf13/cobra"
	"github.com/grafaneus"
	"encoding/json"
	"errors"
)

var PlotDashboardOption string

func InitMetricPlot() *cobra.Command {
	grafana := grafaneus.Grafana{}
	command := cobra.Command{
		Use:   "plot [dashboard]",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			dashboard := args[0]
			metricName := args[1]
			metric, ok := grafaneus.MetricsMetadata[metricName]
			var jsonx string
			if ok {
				jsonx, _ = grafana.GenerateGraph(dashboard, metric.Description, metric.Name)
			} else {
				jsonx, _ = grafana.GenerateGraph(dashboard, metricName, metricName)
			}
			var dash map[string]interface{}
			json.Unmarshal([]byte(jsonx), &dash)
			err := grafana.UploadDashboard(dash)
			if err != nil {
				panic(err)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("two arguments required - 'dashboard' and 'metric'")
			}
			return nil
		},
	}
	return &command
}
