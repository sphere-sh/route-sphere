package route_sphere

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

type Environment struct {
	// Name - unique name of the environment
	Name string `yaml:"name"`

	// Watch - flag to enable/disable watching of the environment
	//
	Watch bool `yaml:"watch"`

	// ConfigurationPath - path to dynamic configuration file(s)
	ConfigurationPath string `yaml:"configuration_path"`

	// Domains - list of domains for the environment
	Domains []Domain `yaml:"domains"`

	// EntryPoints - list of entry points for the environment
	EntryPoints []EntryPoint `yaml:"entryPoints"`

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

	// Watch for changes in the configuration path
	//
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err.Error())
	}
	defer watcher.Close()

	// Handle watcher events
	//
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add the configuration path to the watcher
	//
	err = watcher.Add(e.ConfigurationPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
