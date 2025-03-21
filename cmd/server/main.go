package main

import (
	"log/slog"
	"os"
	route_sphere "route-sphere"
	"sync"
)

var (
	waitGroup = sync.WaitGroup{}
)

func main() {

	// First we verify that the configuration file can be found.
	//
	configurationPath := os.Getenv("ROUTE_SPHERE_CONFIGURATIONS_PATH")
	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		slog.Error("Configuration file not found", "path", configurationPath)
		os.Exit(1)
	}

	// Then, we parse the configuration file into a Configuration struct.
	//
	config, err := route_sphere.ConfigurationFromYamlFile(os.Getenv("ROUTE_SPHERE_CONFIGURATIONS_PATH"))
	if err != nil {
		panic(err)
	}

	// When cloud is enabled start with syncing configurations from the cloud.
	//
	if config.Cloud.Enabled {
		// todo: implement cloud sync
	}

	// Start the Route Sphere server
	//

	// todo: implement route sphere server.

	waitGroup.Wait()
}
