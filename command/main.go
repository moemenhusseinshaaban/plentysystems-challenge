package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "telemetry",
		Short: "Telemetry CLI for logging transactions",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use `telemetry log` to log a transaction")
		},
	}

	rootCmd.AddCommand(GetLogCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
