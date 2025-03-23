package cli_utils

import (
	"context"
	"errors"
	"github.com/alexflint/go-arg"
	cli_commands "route-sphere/cmd/cli/commands"
	"route-sphere/configuration"
)

const ErrInvalidStaticConfiguration = "INVALID_STATIC_CONFIGURATION"

type CliCommandsGroup interface {
	GetCommands() interface{}
}

type CloudCliCommands struct{}

func (c *CloudCliCommands) GetCommands() interface{} {
	var args struct {
		AuthenticationLogin  *cli_commands.AuthenticationLogin  `arg:"subcommand:authentication:login" help:"Login to the cloud provider"`
		AuthenticationLogout *cli_commands.AuthenticationLogout `arg:"subcommand:authentication:logout" help:"Logout from the cloud provider"`

		// Connection commands
		//
		ConnectionList      *cli_commands.ConnectionList      `arg:"subcommand:connection:list" help:"List all connections"`
		ConnectionInstall   *cli_commands.ConnectionInstall   `arg:"subcommand:connection:install" help:"Install a connection"`
		ConnectionUninstall *cli_commands.ConnectionUninstall `arg:"subcommand:connection:uninstall" help:"Uninstall a connection"`
	}
	arg.MustParse(&args)
	return args
}

type LocalCliCommands struct{}

func (c *LocalCliCommands) GetCommands() interface{} {
	var args struct {
	}
	arg.MustParse(&args)
	return args
}

// GetCLICommandGroup - function used to get CLI commands based on the static config.
func GetCLICommandGroup(ctx *context.Context) (CliCommandsGroup, error) {

	// Grab configuration from context
	//
	staticConfiguration := (*ctx).Value("configuration")
	if staticConfiguration == nil {
		return nil, errors.New("CONFIGURATION_NOT_FOUND_IN_CONTEXT")
	}
	c := staticConfiguration.(configuration.StaticConfiguration)

	switch c.Cloud.Enabled {
	case true:
		return &CloudCliCommands{}, nil
	case false:
		return &LocalCliCommands{}, nil
	default:
		return nil, errors.New(ErrInvalidStaticConfiguration)
	}
}
