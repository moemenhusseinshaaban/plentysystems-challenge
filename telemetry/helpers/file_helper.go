// Package helpers contains utility functions for the logging system.
package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenLogFile(filePath string) (*os.File, error) {
	dir := filepath.Dir(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory '%s': %w", dir, err)
		}
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create log file: %w", err)
		}
		return file, nil
	}

	return os.OpenFile(filePath, os.O_RDWR, 0644)
}
