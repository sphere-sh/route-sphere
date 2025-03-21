package cli_commands

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io"
	"log/slog"
	"net/http"
	"os"
	cli_utils "route-sphere/cmd/cli/utils"
)

type ConnectionsList struct{}

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

func (cmd *ConnectionsList) Run(args *ConnectionsList) {

	session, err := cli_utils.SessionGet()
	if err != nil {
		slog.Error("Failed to get session", "error", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/api/v1/connections", nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for _, cookie := range session {
		req.Header.Add("Cookie", cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to get connections", "error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}

	var response V1ConnectionsListResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		slog.Error("Failed to unmarshal response", "error", err)
		os.Exit(1)
	}

	// Create and configure the table
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
	table.Render()

}
