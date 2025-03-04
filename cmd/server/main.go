package main

import (
	"log/slog"
	"os"
	route_sphere "route-sphere"
	"route-sphere/cmd/server/system"
	"sync"
)

var (
	// environmentName - Name of the environment in which the server is running.
	environmentName string
	// environmentsYamlPath - Path pointing to the environments yaml file.
	environmentsYamlPath string
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
		system.EnvironmentsYamlPathValidator,
		system.EnvironmentsNameValidator,
	)
	if err != nil {
		slog.Error("Environment variable validation failed", "error", err.Error())
		os.Exit(1)
	}

	environmentName = os.Getenv("ROUTE_SPHERE_ENVIRONMENT_NAME")
	environmentsYamlPath = os.Getenv("ROUTE_SPHERE_ENVIRONMENTS_PATH")

	waitGroup = sync.WaitGroup{}
}

func main() {

	// Load environments yaml file
	//
	environmentsYaml, _ := os.ReadFile(environmentsYamlPath)
	environments, err := route_sphere.YamlToEnvironments(string(environmentsYaml))
	if err != nil {
		panic(err.Error())
	}

	// Grab the environment matching the current environment name.
	//
	var currentEnvironment route_sphere.Environment
	for _, env := range environments.Items {
		if env.Name == environmentName {
			currentEnvironment = env
			break
		}
	}

	// Set channels on the environment.
	//
	currentEnvironment.ServerUpdateChannel = make(chan route_sphere.ServerUpdate)

	// When the current environment is not found, exit the program
	// with an error message notifying the user that the environment
	// was not found.
	//
	if currentEnvironment.Name == "" {
		slog.Error("environment not found", "name", environmentName)
		os.Exit(1)
	}

	slog.Info("Starting server..")

	// Start watching the environment for configuration changes.
	//
	waitGroup.Add(1)
	go currentEnvironment.StartWatcher(&waitGroup)

	waitGroup.Wait()
}
