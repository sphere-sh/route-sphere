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

func EnvironmentsYamlPathValidator() error {
	environmentsYamlPath := os.Getenv("ROUTE_SPHERE_ENVIRONMENTS_PATH")
	if environmentsYamlPath == "" {
		return errors.New("ROUTE_SPHERE_ENVIRONMENTS_PATH is empty")
	}

	if !regexp.MustCompile("^/").MatchString(environmentsYamlPath) {
		return errors.New("ROUTE_SPHERE_ENVIRONMENTS_PATH must match /")
	}

	if _, err := os.Stat(environmentsYamlPath); os.IsNotExist(err) {
		return errors.New("ROUTE_SPHERE_ENVIRONMENTS_PATH does not exist")
	}

	return nil
}

func EnvironmentsNameValidator() error {
	environmentName := os.Getenv("ROUTE_SPHERE_ENVIRONMENT_NAME")
	if environmentName == "" {
		return errors.New("ROUTE_SPHERE_ENVIRONMENT_NAME is empty")
	}

	if !regexp.MustCompile("^[a-zA-Z]+$").MatchString(environmentName) {
		return errors.New("ROUTE_SPHERE_ENVIRONMENT_NAME is invalid. It should contain only letters")
	}
	return nil
}
