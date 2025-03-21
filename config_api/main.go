package config_api

import (
	"log/slog"
	"sync"
)

// StartConfigurationAPI - starts the configuration REST API.
func StartConfigurationAPI(wg *sync.WaitGroup) {
	defer wg.Done()

	slog.Info("Starting configuration API")
}
