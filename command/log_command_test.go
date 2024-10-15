// command_test.go
package main

import (
	"bytes"
	"telemetry/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogCommand(t *testing.T) {
	// Create a new instance of the log command
	cmd := GetLogCommand()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	logService = service.NewLogService()
	logService.OverrideDriversConfig()

	// Test cases
	tests := []struct {
		name          string
		args          []string
		expectError   bool
		expectedOut   string
		expectedLevel string
	}{
		{
			name:        "missing transaction-id flag",
			args:        []string{"--attribute", "origin=http"},
			expectError: true,
			expectedOut: "Error: required flag(s) \"transaction-id\" not set\n",
		},
		{
			name:        "attributes are optional",
			args:        []string{"--transaction-id", "12345"},
			expectError: false,
		},
		{
			name:        "invalid log level",
			args:        []string{"--transaction-id", "12345", "--attribute", "origin=http", "--log-level", "invalid"},
			expectError: true,
			expectedOut: "Error: failed to log transaction: log-level should be (info, debug, warning, or error)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output.Reset() // Reset output for each test
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, output.String(), tt.expectedOut)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOut, output.String())
			}
		})
	}
}
