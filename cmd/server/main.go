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

	// When the watch flag is set to true, start the
	// environment watcher.
	//
	if currentEnvironment.Watch {
		slog.Info("starting environment watcher", "environment", currentEnvironment.Name)
		waitGroup.Add(1)
		go currentEnvironment.StartWatcher(&waitGroup)
	}

	// Server starting process.
	//
	slog.Info("starting server", "environment", currentEnvironment.Name)

	// Start the entry-point(s)
	//
	for _, entryPoint := range currentEnvironment.EntryPoints {
		// todo: implement entry-point start
		slog.Info("starting entry-point", "name", entryPoint.Name, "address", entryPoint.Address)
	}

	// Open connections to the services
	//
	for _, service := range currentEnvironment.Services {
		// todo: implement service start
		slog.Info("starting service", "service", service)
	}

	// When the entry-points and services are started, we
	// can start connecting them with route rule(s).
	//
	for _, rule := range currentEnvironment.Rules {
		// todo: implement rule start
		slog.Info("starting rule", "rule", rule)
	}

	waitGroup.Wait()
}
