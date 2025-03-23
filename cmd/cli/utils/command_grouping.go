package cli_utils

import (
	"errors"
	cli_commands "route-sphere/cmd/cli/commands"
	"route-sphere/configuration"
)

const ErrInvalidStaticConfiguration = "INVALID_STATIC_CONFIGURATION"

type CliCommandsGroup interface {
	getFeatures() interface{}
}

type CloudCliCommands struct{}

func (c *CloudCliCommands) getFeatures() interface{} {
	return struct {
		AuthenticationLogin *cli_commands.AuthenticationLogin `arg:"subcommand:login" help:"Login to the cloud provider"`
	}{
		AuthenticationLogin: &cli_commands.AuthenticationLogin{},
	}
}

type LocalCliCommands struct{}

func (c *LocalCliCommands) getFeatures() interface{} {
	return struct{}{}
}

// GetCLICommands - function used to get CLI commands based on the static config.
func GetCLICommands(staticConfiguration configuration.StaticConfiguration) (CliCommandsGroup, error) {
	var functionGroups = make(map[string]CliCommandsGroup)

	// Add CLI commands groups here
	//
	functionGroups["cloud"] = &CloudCliCommands{}
	functionGroups["local"] = &LocalCliCommands{}

	// Get CLI commands based on the static configuration
	//
	switch staticConfiguration.Cloud.Enabled {
	case true:
		return functionGroups["cloud"], nil
	case false:
		return functionGroups["local"], nil
	default:
		return nil, errors.New(ErrInvalidStaticConfiguration)
	}
}
