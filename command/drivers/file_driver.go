package commanddrivers

import (
	"fmt"
	"os"
	"telemetry/drivers"
)

type FileDriver struct {
	FilePath string
}

func NewFileDriver(filePath string) drivers.LoggerDriver {
	return &FileDriver{FilePath: filePath}
}

func (f *FileDriver) Log(logEntry drivers.LogEntry) error {
	file, err := os.OpenFile(f.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to log to file: %w", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "[%s] %s: %s | Attributes: %v\n", logEntry.Level, logEntry.TransactionID, logEntry.Timestamp, logEntry.Attributes)

	return nil
}
