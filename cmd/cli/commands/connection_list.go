package cli_commands

import (
	"context"
	"log/slog"
	"os"
)

type ConnectionList struct{}

func (cmd *ConnectionList) Run(args *ConnectionList, ctx *context.Context) {

	// Get session from context.
	//
	session := (*ctx).Value("session")

	// When session is empty, exit with error.
	//
	if session == nil {
		slog.Error("CLI not authenticated, please login using 'route-sphere authentication:login' command")
		os.Exit(1)
	}

}
