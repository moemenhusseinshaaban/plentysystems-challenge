package service

import (
	"fmt"
	"sync"

	"telemetry/config"
	"telemetry/drivers"
)

type LogService struct {
	drivers []drivers.LoggerDriver
}

func NewLogService() *LogService {
	defaultDrivers := config.LoadDefaultDriversConfig()
	logService := LogService{drivers: []drivers.LoggerDriver{}}

	logService.drivers = append(logService.drivers, defaultDrivers...)

	return &logService
}

func (s *LogService) OverrideDriversConfig(loggerDrivers ...drivers.LoggerDriver) {
	s.drivers = []drivers.LoggerDriver{}
	s.drivers = append(s.drivers, loggerDrivers...)
}

func (s *LogService) AppendDriversConfig(loggerDrivers ...drivers.LoggerDriver) {
	s.drivers = append(s.drivers, loggerDrivers...)
}

func (s *LogService) Log(level drivers.LogLevel, transactionID string, attributes map[string]string) error {
	logEntry := drivers.NewLogEntry(level, transactionID, attributes)

	return logTransaction(*logEntry, s.drivers)
}

func (s *LogService) LogInfo(transactionID string, attributes map[string]string) error {
	return s.Log(drivers.Info, transactionID, attributes)
}

func (s *LogService) LogDebug(transactionID string, attributes map[string]string) error {
	return s.Log(drivers.Debug, transactionID, attributes)
}

func (s *LogService) LogWarning(transactionID string, attributes map[string]string) error {
	return s.Log(drivers.Warning, transactionID, attributes)
}

func (s *LogService) LogError(transactionID string, attributes map[string]string) error {
	return s.Log(drivers.Error, transactionID, attributes)
}

func logTransaction(logEntry drivers.LogEntry, loggerDrivers []drivers.LoggerDriver) error {
	var wg sync.WaitGroup
	errorChannel := make(chan error, len(loggerDrivers))

	for _, driver := range loggerDrivers {
		wg.Add(1)
		go func(driver drivers.LoggerDriver) {
			defer wg.Done()
			if err := driver.Log(logEntry); err != nil {
				errorChannel <- err
			}
		}(driver)
	}

	wg.Wait()
	close(errorChannel)

	// Collect errors
	var combinedError error
	for err := range errorChannel {
		combinedError = fmt.Errorf("%v | %w", combinedError, err)
	}

	return combinedError
}
