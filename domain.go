package route_sphere

type Domain struct {
	// URL - domain URL
	URL string `yaml:"url"`

	// TLS - TLS Configuration
	TLS struct {
		// Cert - path to certificate file
		Cert string `yaml:"cert"`
		// Key - path to key file
		Key string `yaml:"key"`
	}
}
