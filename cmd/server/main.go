package main

import (
	"log/slog"
	"os"
	route_sphere "route-sphere"
	"sync"
	"time"
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
	_, err := route_sphere.ConfigurationFromYamlFile(os.Getenv("ROUTE_SPHERE_CONFIGURATIONS_PATH"))
	if err != nil {
		panic(err)
	}

	// Keep process running
	//
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.Sleep(10000 * time.Minute)

	wg.Wait()

}
