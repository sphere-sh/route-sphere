package main

import (
	"context"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"reflect"
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
}

func main() {
	defer logFile.Close()

	// Construct context of the application.
	//
	ctx := context.Background()

	// When we have a session we add it to the context. Otherwise, we put an
	// empty string in the context indicating that we don't have a session.
	//
	if cli_utils.HasSession() {
		sessionString, err := cli_utils.GetSession()
		if err != nil {
			slog.Error("Failed to get session", "error", err)
			os.Exit(1)
		}
		ctx = context.WithValue(ctx, "session", sessionString)
	} else {
		ctx = context.WithValue(ctx, "session", "")
	}

	// Load static configuration and put it in the context.
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
	ctx = context.WithValue(ctx, "configuration", staticConfiguration)

	// Get CLI features for the current context.
	//
	commandGroup, err := cli_utils.GetCLICommandGroup(&ctx)
	commands := commandGroup.GetCommands()

	// Find and execute the selected command
	//
	cmdValue := reflect.ValueOf(commands)
	cmdType := cmdValue.Type()

	// Loop through fields to find which command was selected (non-nil)
	for i := 0; i < cmdValue.NumField(); i++ {
		fieldValue := cmdValue.Field(i)
		fieldName := cmdType.Field(i).Name

		// Check if this field is a pointer and not nil
		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			slog.Info("Executing command", "command", fieldName)

			// Call Run() method if it exists
			method := fieldValue.MethodByName("Run")
			if method.IsValid() {
				// Create argument list - pass the command itself as the argument
				args := []reflect.Value{fieldValue}
				args = append(args, reflect.ValueOf(&ctx))

				method.Call(args)
			} else {
				slog.Error("Command does not implement Run method", "command", fieldName)
				os.Exit(1)
			}

			// Exit after executing the command
			return
		}
	}

	slog.Error("No command selected")
	os.Exit(1)
}
