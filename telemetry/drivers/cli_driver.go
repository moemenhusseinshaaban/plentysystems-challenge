// Package drivers contains the CLI logger implementation.
package drivers

import (
	"log"
)

// CliDriver is a logger driver that outputs logs to the CLI.
type CliDriver struct{}

// NewCliDriver creates a new instance of CliDriver.
func NewCliDriver() LoggerDriver {
	return &CliDriver{}
}

// Log writes the log message to the CLI.
func (c *CliDriver) Log(logEntry LogEntry) error {
	log.Printf(
		"[%s] %s: %s | Attributes: %v \n",
		logEntry.Level,
		logEntry.TransactionID,
		logEntry.Timestamp,
		logEntry.Attributes,
	)

	return nil
}
