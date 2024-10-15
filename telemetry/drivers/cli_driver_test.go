package drivers

import (
	"bytes"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCliDriver_Log(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil) // Restore original log output

	driver := NewCliDriver()
	logEntry := LogEntry{
		Level: Info,
		Transaction: Transaction{
			TransactionID: "123",
			Attributes:    map[string]string{"origin": "http"},
		},
		Timestamp: time.Now(),
	}

	driver.Log(logEntry)

	// Check that the log contains the expected content
	output := buf.String()
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "123")
	assert.Contains(t, output, "origin")
	assert.Contains(t, output, "http")
}
