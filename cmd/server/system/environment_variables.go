package system

import (
	"errors"
	"os"
	"regexp"
)

type EnvironmentVariableValidator func() error

func EnvironmentVariableValidate(validators ...EnvironmentVariableValidator) error {
	for _, validator := range validators {
		if err := validator(); err != nil {
			return err
		}
	}
	return nil
}

func ConfigurationsYamlPathValidator() error {
	configurationsYamlPath := os.Getenv("ROUTE_SPHERE_CONFIGURATIONS_PATH")
	if configurationsYamlPath == "" {
		return errors.New("ROUTE_SPHERE_CONFIGURATIONS_PATH is empty")
	}

	if !regexp.MustCompile("^/").MatchString(configurationsYamlPath) {
		return errors.New("ROUTE_SPHERE_CONFIGURATIONS_PATH must match /")
	}

	if _, err := os.Stat(configurationsYamlPath); os.IsNotExist(err) {
		return errors.New("ROUTE_SPHERE_CONFIGURATIONS_PATH does not exist")
	}

	return nil
}

func ConfigurationNameValidator() error {
	configurationName := os.Getenv("ROUTE_SPHERE_CONFIGURATION_NAME")
	if configurationName == "" {
		return errors.New("ROUTE_SPHERE_CONFIGURATION_NAME is empty")
	}

	if !regexp.MustCompile("^[a-zA-Z]+$").MatchString(configurationName) {
		return errors.New("ROUTE_SPHERE_CONFIGURATION_NAME is invalid. It should contain only letters")
	}
	return nil
}
