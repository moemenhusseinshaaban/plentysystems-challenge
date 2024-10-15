package drivers

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"telemetry/helpers"

	"github.com/stretchr/testify/assert"
)

func TestJsonDriver_Log(t *testing.T) {
	filePath := "test_logs.json"
	defer os.Remove(filePath)

	driver := NewJsonDriver(filePath)
	logEntry := LogEntry{
		Level: Info,
		Transaction: Transaction{
			TransactionID: "123",
			Attributes:    map[string]string{"origin": "http"},
		},
		Timestamp: time.Now(),
	}

	err := driver.Log(logEntry)
	assert.NoError(t, err)

	// Verify file contents
	file, err := helpers.OpenLogFile(filePath)
	assert.NoError(t, err)
	defer file.Close()

	var logEntries []LogEntry
	err = json.NewDecoder(file).Decode(&logEntries)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logEntries))
	assert.Equal(t, "123", logEntries[0].TransactionID)
	assert.Equal(t, "http", logEntries[0].Attributes["origin"])
}
