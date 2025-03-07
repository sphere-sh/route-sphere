package route_sphere

type Configuration struct {
	// Name - unique name of the configuration
	Name string `yaml:"name"`

	// Providers - List of configuration providers
	Providers []Provider `yaml:"providers"`

	// EntryPoints - List of entry points
	EntryPoints []EntryPoint `yaml:"entryPoints"`

	// UpdateChannel - channel to send configuration updates
	UpdateChannel chan ConfigurationUpdate `yaml:"-"`
}
