// Package config provides configurations for logging drivers.
package config

import (
	"telemetry/drivers"
)

// LoadDefaultDriversConfig returns the default configuration for logger drivers.
func LoadDefaultDriversConfig() []drivers.LoggerDriver {
	return []drivers.LoggerDriver{
		drivers.NewCliDriver(),
		drivers.NewJsonDriver("logs/app.json"),
	}
}
