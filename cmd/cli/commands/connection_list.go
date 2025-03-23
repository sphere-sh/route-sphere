package cli_commands

import (
	"context"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type V1ConnectionsListResponse struct {
	Data struct {
		Connections []struct {
			Id          string `json:"id"`
			State       string `json:"state"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Tls_state   string `json:"tls_state"`
		} `json:"connections"`
	} `json:"data"`
}

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
	sessionP := []string{}
	err = json.Unmarshal([]byte(sessionString), &sessionP)

	for _, cookie := range sessionP {
		req.Header.Add("Cookie", cookie)
	}

	// Send request
	//
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send HTTP request", "error", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Read response body
	//
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}

	// Render responseBody->data->connections as string
	//
	var response V1ConnectionsListResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		slog.Error("Failed to unmarshal response", "error", err)
		os.Exit(1)
	}

	// Create and configure the table
	//
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "State", "Name", "Description", "TLS State"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	// Add data to the table
	//
	for _, connection := range response.Data.Connections {
		table.Append([]string{
			connection.Id,
			connection.State,
			connection.Name,
			connection.Description,
			connection.Tls_state,
		})
	}

	// Render the table
	//
	table.Render()

	// Close the response body
	//
	os.Exit(0)
}
