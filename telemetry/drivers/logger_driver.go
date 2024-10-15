// Package drivers provides various logging drivers.
package drivers

import "time"

// LogLevel defines the level of the log message.
type LogLevel string

const (
	Debug   LogLevel = "DEBUG"
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARNING"
	Error   LogLevel = "ERROR"
)

// Transaction represents the transaction data
type Transaction struct {
	TransactionID string
	Attributes    map[string]string
}

// LogEntry represents the log entry details
type LogEntry struct {
	Level LogLevel
	Transaction
	Timestamp time.Time
}

func NewLogEntry(level LogLevel, transactionID string, attributes map[string]string) *LogEntry {
	return &LogEntry{
		Level: level,
		Transaction: Transaction{
			TransactionID: transactionID,
			Attributes:    attributes,
		},
		Timestamp: time.Now(),
	}
}

// LoggerDriver is an interface for logging drivers.
type LoggerDriver interface {
	Log(logEntry LogEntry) error
}
