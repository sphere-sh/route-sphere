package main

import (
	"log/slog"
	"os"
	route_sphere "route-sphere"
	"route-sphere/cmd/server/system"
	"sync"
)

var (
	// configurationsYamlPath - Path to the configurations yaml file.
	configurationsYamlPath string
	// configurationName - Name of the configuration to use.
	configurationName string
	// waitGroup - Wait group to wait for all goroutines to finish.
	waitGroup sync.WaitGroup
)

func init() {

	// -----------------------------------------------------------------------
	//
	// Validation of environment variables
	//
	// Before we can start the server, we want to make sure that the environment
	// variables are safe to use. Environment variables are safe to use when;
	// 	1. Edge cases are handled,
	// 	2. No malicious input is provided,
	//
	// -----------------------------------------------------------------------

	err := system.EnvironmentVariableValidate(
		system.ConfigurationsYamlPathValidator,
		system.ConfigurationNameValidator,
	)
	if err != nil {
		slog.Error("Environment variable validation failed", "error", err.Error())
		os.Exit(1)
	}

	configurationsYamlPath = os.Getenv("ROUTE_SPHERE_CONFIGURATIONS_PATH")
	configurationName = os.Getenv("ROUTE_SPHERE_CONFIGURATION_NAME")

	waitGroup = sync.WaitGroup{}
}

func main() {

	// Load configurations structure
	//
	configurationsYaml, _ := os.ReadFile(configurationsYamlPath)
	configurations, err := route_sphere.ConfigurationYamlToStruct(string(configurationsYaml))
	if err != nil {
		slog.Error("Error parsing configurations", "error", err.Error())
		os.Exit(1)
	}

	// Grab the configuration matching the current configuration name.
	//
	var currentConfiguration route_sphere.Configuration
	for _, config := range configurations.Items {
		if config.Name == configurationName {
			currentConfiguration = *config
			break
		}
	}

	// When the current configuration is not found, exit the program
	//
	if currentConfiguration.Name == "" {
		slog.Error("Could not find configuration with the provided name.", "name", configurationName)
		os.Exit(1)
	}

	// Before we can start the providers we need to create a channel which
	// the providers can use to notify the configuration of updates.
	//
	currentConfiguration.UpdateChannel = make(chan route_sphere.ConfigurationUpdate, 100)

	// Start the listener for configuration updates
	//
	waitGroup.Add(1)
	go currentConfiguration.ListenForConfigurationUpdates(&waitGroup)

	// Start configuration watcher if `configuration.watch` is set the
	// boolean value `true`.
	//
	for _, provider := range currentConfiguration.Providers {
		if provider.Watch {
			waitGroup.Add(1)
			go provider.StartWatcher(&waitGroup, &currentConfiguration.UpdateChannel)
		}
	}

	// Start entryPoints defined in the configuration
	//
	for _, entryPoint := range currentConfiguration.EntryPoints {
		waitGroup.Add(1)
		go entryPoint.Start(&waitGroup)
	}

	// Server initialization process.
	//
	slog.Info("starting server", "configuration", currentConfiguration.Name)

	waitGroup.Wait()
}
