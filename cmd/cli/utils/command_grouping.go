package cli_utils

import (
	"errors"
	"github.com/alexflint/go-arg"
	cli_commands "route-sphere/cmd/cli/commands"
	"route-sphere/configuration"
)

const ErrInvalidStaticConfiguration = "INVALID_STATIC_CONFIGURATION"

type CliCommandsGroup interface {
	GetFeatures() interface{}
}

type CloudCliCommands struct{}

func (c *CloudCliCommands) GetFeatures() interface{} {
	var args struct {
		AuthenticationLogin  *cli_commands.AuthenticationLogin  `arg:"subcommand:authentication:login" help:"Login to the cloud provider"`
		AuthenticationLogout *cli_commands.AuthenticationLogout `arg:"subcommand:authentication:logout" help:"Logout from the cloud provider"`
	}
	arg.MustParse(&args)
	return args
}

type LocalCliCommands struct{}

func (c *LocalCliCommands) GetFeatures() interface{} {
	var args struct {
	}
	arg.MustParse(&args)
	return args
}

// GetCLICommandGroup - function used to get CLI commands based on the static config.
func GetCLICommandGroup(staticConfiguration configuration.StaticConfiguration) (CliCommandsGroup, error) {
	switch staticConfiguration.Cloud.Enabled {
	case true:
		return &CloudCliCommands{}, nil
	case false:
		return &LocalCliCommands{}, nil
	default:
		return nil, errors.New(ErrInvalidStaticConfiguration)
	}
}
