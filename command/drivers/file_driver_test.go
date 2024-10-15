package commanddrivers

import (
	"fmt"
	"os"
	"telemetry/drivers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileDriver_Log(t *testing.T) {
	filePath := "test_logs.txt"
	defer os.Remove(filePath)

	driver := NewFileDriver(filePath)

	logEntry := drivers.LogEntry{
		Level: drivers.Info,
		Transaction: drivers.Transaction{
			TransactionID: "123",
			Attributes:    map[string]string{"origin": "http"},
		},
		Timestamp: time.Now(),
	}

	err := driver.Log(logEntry)
	assert.NoError(t, err)

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("failed to read log file: %v", err)
	}

	expectedLogLine := fmt.Sprintf(
		"[%s] %s: %s | Attributes: %v\n",
		logEntry.Level,
		logEntry.TransactionID,
		logEntry.Timestamp,
		logEntry.Attributes,
	)
	if string(content) != expectedLogLine {
		t.Errorf("unexpected log content: got %q, want %q", string(content), expectedLogLine)
	}
}
