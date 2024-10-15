package main

import (
	"errors"
	"fmt"
	"strings"
	"telemetry/drivers"
	"telemetry/service"

	"github.com/spf13/cobra"
	commanddrivers "main.go/drivers"
)

var (
	transactionID string
	attributes    []string
	logLevel      string
	attributeMaps map[string]string
	logService    *service.LogService
)

func init() {
	fmt.Println("Initializing commands...")
}

// GetLogCommand returns the log command with its configuration.
func GetLogCommand() *cobra.Command {
	logCmd := &cobra.Command{
		Use:   "log",
		Short: "Log a transaction",
		RunE:  runLogCommand,
	}

	logCmd.Flags().StringVarP(&transactionID, "transaction-id", "t", "", "Transaction ID (required)")
	logCmd.Flags().StringArrayVarP(&attributes, "attribute", "a", []string{}, "Transaction attributes (key=value)")
	logCmd.Flags().StringVarP(&logLevel, "log-level", "l", "info", "Log level (info, debug, warning, error)")

	logCmd.MarkFlagRequired("transaction-id")

	logService = service.NewLogService()

	driversList := []drivers.LoggerDriver{
		commanddrivers.NewFileDriver("logs/app.txt"),
	}
	logService.AppendDriversConfig(driversList...)

	return logCmd
}

func runLogCommand(cmd *cobra.Command, args []string) error {
	attributeMaps = parseAttributes(attributes)

	if err := logTransaction(); err != nil {
		return fmt.Errorf("failed to log transaction: %w", err)
	}

	fmt.Println("Transaction logged successfully")
	return nil
}

func parseAttributes(attrs []string) map[string]string {
	attributesMap := make(map[string]string)
	for _, attr := range attrs {
		if key, value := splitKeyValue(attr); key != "" {
			attributesMap[key] = value
		}
	}

	return attributesMap
}

func splitKeyValue(input string) (key string, value string) {
	if parts := strings.SplitN(input, "=", 2); len(parts) == 2 {
		key = parts[0]
		value = parts[1]
	}

	return
}

func logTransaction() error {
	switch strings.ToLower(logLevel) {
	case "info":
		return logService.LogInfo(transactionID, attributeMaps)
	case "debug":
		return logService.LogDebug(transactionID, attributeMaps)
	case "warning":
		return logService.LogWarning(transactionID, attributeMaps)
	case "error":
		return logService.LogError(transactionID, attributeMaps)
	default:
		return errors.New("log-level should be (info, debug, warning, or error)")
	}
}
