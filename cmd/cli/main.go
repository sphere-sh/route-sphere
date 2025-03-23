package main

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	cli_utils "route-sphere/cmd/cli/utils"
	"route-sphere/configuration"
)

var (
	// Storage related paths
	//
	routeSpherePath = "/etc/route-sphere"

	// Configuration
	//
	staticConfigurationPath = "/etc/route-sphere/route-sphere.yaml"

	// Logging
	//
	logFile *os.File
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func init() {
	var err error

	// Setup paths for route sphere and static configuration file
	//
	routeSpherePath = getEnv("ROUTE_SPHERE_PATH", routeSpherePath)
	if _, err := os.Stat(routeSpherePath); os.IsNotExist(err) {
		slog.Error("Route Sphere path not found", "path", routeSpherePath)
		os.Exit(1)
	}

	staticConfigurationPath = getEnv("ROUTE_SPHERE_CONFIG_PATH", staticConfigurationPath)
	if _, err := os.Stat(staticConfigurationPath); os.IsNotExist(err) {
		slog.Error("Static configuration file not found", "path", staticConfigurationPath)
		os.Exit(1)
	}

	// Create CLI path
	//
	err = os.MkdirAll(routeSpherePath+"/cli", 0755)
	if err != nil {
		panic(err)
	}

	// Create storage path for CLI logs.
	//
	os.MkdirAll(routeSpherePath+"/cli/logs", 0755)
	if err != nil {
		panic(err)
	}

	// Setup logging
	//
	logFile, err = os.OpenFile(routeSpherePath+"/cli/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	var logLevel = new(slog.LevelVar)

	logger := slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(logger))
}

func main() {
	defer logFile.Close()

	// Load static configuration
	//
	var staticConfiguration configuration.StaticConfiguration

	fileContent, err := os.ReadFile(staticConfigurationPath)
	if err != nil {
		slog.Error("Failed to read static configuration file", "error", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(fileContent, &staticConfiguration)
	if err != nil {
		slog.Error("Failed to unmarshal static configuration file", "error", err)
		os.Exit(1)
	}

	// Get CLI features
	//
	commandGroup, err := cli_utils.GetCLICommandGroup(staticConfiguration)
	if err != nil {
		slog.Error("Failed to get CLI features", "error", err)
		os.Exit(1)
	}

	commands := commandGroup.GetCommands()

	if commands == nil {
		slog.Error("CLI features not found")
		os.Exit(1)
	}

}
