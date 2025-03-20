package route_sphere

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Configuration struct {
	Cloud struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"cloud"`
}

// ConfigurationFromYamlFile - reads content from a yaml file and returns a Configuration struct
func ConfigurationFromYamlFile(path string) (*Configuration, error) {

	// Read file
	//
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal yaml
	//
	var config Configuration
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
