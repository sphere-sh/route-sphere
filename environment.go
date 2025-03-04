package route_sphere

import "sync"

type Environment struct {
	// Name - unique name of the environment
	Name string `yaml:"name"`

	// EntryPoints - list of entry points for the environment
	EntryPoints []string `yaml:"entryPoints"`

	// Rules - list of route rules for the environment
	Rules []string `yaml:"rules"`

	// Services - list of services for the environment
	Services []string `yaml:"services"`

	// ServerUpdateChannel - channel to send server updates
	ServerUpdateChannel chan ServerUpdate `yaml:"-"`
}

type ServerUpdate struct{}

// StartWatcher - Starts the environment watcher
func (e Environment) StartWatcher(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
}

// EntryPointsUpdate - Updates the entry points for the environment
func (e Environment) EntryPointsUpdate() {}

// RulesUpdate - Updates the rules for the environment
func (e Environment) RulesUpdate() {}

// ServicesUpdate - Updates the services for the environment
func (e Environment) ServicesUpdate() {}
