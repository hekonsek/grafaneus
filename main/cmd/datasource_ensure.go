package cmd

import (
	"github.com/spf13/cobra"
	"github.com/grafaneus"
	"fmt"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "datasource-ensure",
		Short: "List metrics available in Prometheus.",
		Long:  `List metrics available in Prometheus. Include optional metadata about the most commonly used metrics.`,
		Run: func(cmd *cobra.Command, args []string) {
			dataSources, _ := grafaneus.ListDataSources()
			if len(dataSources) > 0 {
				fmt.Println("Datasource exists.")
			} else {
				fmt.Println("Datasource doesn't exist. Creating...")
				grafaneus.CreateDataSource()
			}
		},
	})
}