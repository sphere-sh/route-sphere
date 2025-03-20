package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"log/slog"
	"net/http"
	"os"
	cli_commands "route-sphere/cmd/cli/commands"
)

func init() {
	err := os.MkdirAll("/etc/route-sphere/cli", 0755)
	if err != nil {
		panic(err)
	}
}

func main() {

	var args struct {
		AuthenticationLogin  *cli_commands.AuthenticationLogin  `arg:"subcommand:authentication:login"`
		AuthenticationLogout *cli_commands.AuthenticationLogout `arg:"subcommand:authentication:logout"`

		// Connections related commands
		//
		ConnectionsList *cli_commands.ConnectionsList `arg:"subcommand:connections:list"`
	}

	arg.MustParse(&args)

	if args.AuthenticationLogin != nil {
		args.AuthenticationLogin.Run(args.AuthenticationLogin)
	}

	// Check if `session` file exists in the /etc/route-sphere/cli directory
	//
	_, err := os.Stat("/etc/route-sphere/cli/session")
	if os.IsNotExist(err) {
		slog.Error("Please authenticate using `route-sphere authentication:login`")
		os.Exit(1)
	}

	// Verify that the session is valid
	//
	err = sessionCheck()
	if err != nil {
		slog.Error("Failed to authenticate", "error", err)
		os.Exit(1)
	}

	if args.AuthenticationLogout != nil {
		args.AuthenticationLogout.Run(args.AuthenticationLogout)
	}

}

func sessionCheck() error {
	sessionFile, err := os.ReadFile("/etc/route-sphere/cli/session")
	if err != nil {
		return fmt.Errorf("Failed to read session file: %w", err)
	}

	session := []string{}
	err = json.Unmarshal(sessionFile, &session)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal session file: %w", err)
	}

	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/user", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for _, cookie := range session {
		req.Header.Add("Cookie", cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("session check failed: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to authenticate: %d", resp.StatusCode)
	}

	return nil
}
