package cli_commands

import (
	"context"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io"
	"log/slog"
	"os"
	"route-sphere/cmd/cli/cloud"
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

	cloudClient := (*ctx).Value("cloudClient").(*cloud.CloudHTTPClient)

	// Collect all connections from the Route Sphere cloud.
	//
	resp, err := cloudClient.Get("/api/v1/connections")
	if err != nil {
		slog.Error("Failed to create HTTP request", "error", err)
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
