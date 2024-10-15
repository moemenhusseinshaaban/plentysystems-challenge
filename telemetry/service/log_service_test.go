package service

import (
	"errors"
	"testing"

	"telemetry/drivers"

	"github.com/stretchr/testify/assert"
)

// MockDriver is a mock implementation of LoggerDriver for testing.
type MockDriver struct {
	LoggedEntries []drivers.LogEntry
	ShouldFail    bool
}

func (m *MockDriver) Log(logEntry drivers.LogEntry) error {
	if m.ShouldFail {
		return errors.New("mock error")
	}
	m.LoggedEntries = append(m.LoggedEntries, logEntry)

	return nil
}

func TestLogService_Log(t *testing.T) {
	driver1 := &MockDriver{}
	driver2 := &MockDriver{}
	logService := &LogService{
		drivers: []drivers.LoggerDriver{driver1, driver2},
	}

	transactionID := "123"
	attributes := map[string]string{"origin": "http"}
	err := logService.Log(drivers.Info, transactionID, attributes)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(driver1.LoggedEntries))
	assert.Equal(t, 1, len(driver2.LoggedEntries))
	assert.Equal(t, transactionID, driver1.LoggedEntries[0].TransactionID)
	assert.Equal(t, drivers.Info, driver1.LoggedEntries[0].Level)
}

func TestLogService_Log_WithError(t *testing.T) {
	driver1 := &MockDriver{}
	driver2 := &MockDriver{ShouldFail: true}
	logService := &LogService{
		drivers: []drivers.LoggerDriver{driver1, driver2},
	}

	transactionID := "456"
	attributes := map[string]string{"origin": "http"}
	err := logService.Log(drivers.Error, transactionID, attributes)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock error")
	assert.Equal(t, 1, len(driver1.LoggedEntries))
	assert.Equal(t, 0, len(driver2.LoggedEntries)) // Failed driver should not log
}

func TestLogService_OverrideDriversConfig(t *testing.T) {
	driver1 := &MockDriver{}
	driver2 := &MockDriver{}
	logService := NewLogService()
	logService.OverrideDriversConfig(driver1, driver2)

	assert.Equal(t, 2, len(logService.drivers))
	assert.Equal(t, driver1, logService.drivers[0])
	assert.Equal(t, driver2, logService.drivers[1])
}

func TestLogService_AppendDriversConfig(t *testing.T) {
	driver1 := &MockDriver{}
	driver2 := &MockDriver{}
	logService := NewLogService()
	logService.AppendDriversConfig(driver1)
	logService.AppendDriversConfig(driver2)

	assert.Equal(t, 4, len(logService.drivers))     // Including the default drivers
	assert.Equal(t, driver1, logService.drivers[2]) // Third appended driver
	assert.Equal(t, driver2, logService.drivers[3]) // Fourth appended driver
}
