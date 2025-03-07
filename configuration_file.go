package route_sphere

type ConfigurationFile struct {
	Domains  []Domain  `yaml:"domains"`
	Routes   []Route   `yaml:"routes"`
	Services []Service `yaml:"services"`
}

type Route struct{}
type Service struct{}
