package route_sphere

import (
	"errors"
	"gopkg.in/yaml.v3"
)

const (
	// ErrEnvironmentNameNotUnique - Error message for non-unique environment name.
	ErrEnvironmentNameNotUnique = "ENVIRONMENT_NAME_NOT_UNIQUE"
)

type Environments struct {
	Items []Environment `yaml:"environments"`
}

// YamlToEnvironments - Convert yaml string to Environments struct
func YamlToEnvironments(yamlString string) (Environments, error) {
	var environments Environments
	err := yaml.Unmarshal([]byte(yamlString), &environments)
	if err != nil {
		return Environments{}, err
	}
	return environments, nil
}

// ToYaml - Converts Environments struct to yaml string
func (e *Environments) ToYaml() (string, error) {
	out, err := yaml.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Exists - Check if there is an environment with the given name
func (e *Environments) Exists(name string) bool {
	for _, env := range e.Items {
		if env.Name == name {
			return true
		}
	}
	return false
}

// Add - Add an environment to the list
func (e *Environments) Add(env Environment) error {
	if e.Exists(env.Name) {
		return errors.New(ErrEnvironmentNameNotUnique)
	}

	return nil
}
