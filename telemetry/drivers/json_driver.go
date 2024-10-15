// Package drivers contains the JSON logger implementation.
package drivers

import (
	"encoding/json"
	"fmt"
	"io"

	"telemetry/helpers"
)

// JsonDriver is a logger driver that outputs logs to a JSON file.
type JsonDriver struct {
	FilePath string
}

// NewJsonDriver creates a new instance of JsonDriver.
func NewJsonDriver(filePath string) LoggerDriver {
	return &JsonDriver{FilePath: filePath}
}

// Log writes the log message to a JSON file.
func (j *JsonDriver) Log(logEntry LogEntry) error {
	file, err := helpers.OpenLogFile(j.FilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var logEntries []LogEntry
	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	if len(fileData) > 0 {
		if err := json.Unmarshal(fileData, &logEntries); err != nil {
			return fmt.Errorf("error unmarshaling existing logs: %w", err)
		}
	}

	logEntries = append(logEntries, logEntry)

	updatedData, err := json.MarshalIndent(logEntries, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling updated logs: %w", err)
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating file: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking file: %w", err)
	}

	if _, err := file.Write(updatedData); err != nil {
		return fmt.Errorf("error writing updated logs: %w", err)
	}

	return nil
}
