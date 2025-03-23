package cli_commands

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

type ConnectionList struct{}

func (cmd *ConnectionList) Run(args *ConnectionList, ctx *context.Context) {

	// Get session from context.
	//
	session := (*ctx).Value("session")
	if session == nil {
		slog.Error("CLI not authenticated, please login using 'route-sphere authentication:login' command")
		os.Exit(1)
	}

	// Collect all connections from the Route Sphere cloud.
	//
	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/api/v1/connections", nil)
	if err != nil {
		slog.Error("Failed to create HTTP request", "error", err)
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add session to the request as cookie
	//
	sessionString := session.(string)
	for _, cookie := range sessionString {
		req.Header.Add("Cookie", string(cookie))
	}

	// Send request
	//
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send HTTP request", "error", err)
		os.Exit(1)
	}

	println(resp.Status)
}
