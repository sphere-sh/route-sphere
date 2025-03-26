package cli_commands

import (
	"context"
	"log/slog"
	"os"
)

type ConnectionInstall struct {
	Id string `arg:"positional,required" help:"Connection ID"`
}

func (cmd *ConnectionInstall) Run(args *ConnectionInstall, ctx *context.Context) {

	// Before we start the installation process, we must check if we have
	// a session available otherwise we cannot proceed.
	//
	session := (*ctx).Value("session")
	if session == nil || session == "" {
		slog.Error("CLI not authenticated, please login using 'route-sphere authentication:login' command")
		os.Exit(1)
	}

	slog.Info("Installing connection", "id", args.Id)

	// First, we have to make sure that the connection exists in the cloud.
	//
	existenceCheckError := checkConnectionExistence(args.Id)
	if existenceCheckError != nil {
		slog.Error("Error occurred while checking connection existence", "error", existenceCheckError)
		os.Exit(1)
	}

	// Second, we have to make sure that there is no other process installing the connection.
	//
	otherProcessInstallingConnection := checkOtherProcessInstallingConnection(args.Id)
	if otherProcessInstallingConnection != nil {
		slog.Error("Error occurred while checking other process installing connection", "error", otherProcessInstallingConnection)
		os.Exit(1)
	}

	// Third, we have to make sure that there is no other connection installed.
	//
	otherConnectionInstalled := checkOtherConnectionInstalled()
	if otherConnectionInstalled != nil {
		slog.Error("Error occurred while checking other connection installed", "error", otherConnectionInstalled)
		os.Exit(1)
	}

	// Finally, we execute the installation process.
	//
	connectionBeingInstalledMarkingError := markConnectionBeingInstalled(args.Id)
	if connectionBeingInstalledMarkingError != nil {
		slog.Error("Error occurred while marking connection being installed", "error", connectionBeingInstalledMarkingError)
		os.Exit(1)
	}

}

// checkConnectionExistence - when no error is returned the connection exists.
func checkConnectionExistence(id string) error {
	return nil
}

// checkOtherProcessInstallingConnection - when no error is returned there is no other process installing the connection.
func checkOtherProcessInstallingConnection(id string) error {
	return nil
}

// checkOtherConnectionInstalled - when no error is returned there is no other connection installed.
func checkOtherConnectionInstalled() error {
	return nil
}

// markConnectionBeingInstalled - mark that there is a process being installed. This marking is stored locally
// and on the cloud.
func markConnectionBeingInstalled(id string) error {
	return nil
}
