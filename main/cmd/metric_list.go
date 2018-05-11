package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/hekonsek/grafaneus"
)

func InitMetricList() *cobra.Command {
	command := cobra.Command{
		Use:   "list",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Name\t\t\t\t\t\tType\t\tDescription")
			fmt.Println("=============================================================================================================")
			metrics, err := grafaneus.ListMetrics()
			if err != nil {
				panic(err)
			}
			for _, metric:= range metrics {
				fmt.Printf("%s\t\t\t\t\t%s\t\t%s\n", metric.Name, metric.Type, metric.Description)
			}
		},
	}
	return &command
}