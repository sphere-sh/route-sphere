package route_sphere

import "gopkg.in/yaml.v3"

type Configurations struct {
	// Items - list of configurations
	Items []*Configuration `yaml:"configurations"`
}

// ConfigurationYamlToStruct - Convert yaml string to Configurations struct
func ConfigurationYamlToStruct(yamlString string) (Configurations, error) {
	var configurations Configurations
	err := yaml.Unmarshal([]byte(yamlString), &configurations)
	if err != nil {
		return Configurations{}, err
	}
	return configurations, nil
}
