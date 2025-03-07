package route_sphere

import (
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Provider struct {
	// Type - type of the provider
	Type string `yaml:"type"`

	// Watch - flag to enable/disable watching of the provider
	Watch bool `yaml:"watch"`

	// Directory - directory to watch for configuration changes when type equals `directory`
	Directory string `yaml:"directory"`
}

// StartWatcher - watch the provider for configuration changes
func (p *Provider) StartWatcher(wg *sync.WaitGroup, updateChannel *chan ConfigurationUpdate) {
	defer wg.Done()
	defer close(*updateChannel)

	switch p.Type {
	case "directory":
		p.initDirectory(updateChannel)
		p.watchDirectory(updateChannel)
	default:
		panic("UNKNOWN_PROVIDER_TYPE")
	}

}

// initDirectory - initialize the directory provider
func (p *Provider) initDirectory(updateChannel *chan ConfigurationUpdate) {
	// todo: implement initial load.

	// Scan `p.Directory` for configuration files and load them
	// into the configuration update channel.
	//
	files, err := os.ReadDir(p.Directory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(p.Directory, file.Name())

		yamlFile, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		// Parse YAML to ConfigurationFile.
		//
		var config ConfigurationFile
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(err)
		}

		// Send configuration update to the channel
		//
		*updateChannel <- ConfigurationUpdate{
			Action: "ADD",
			Key:    "domains",
			Value:  config.Domains,
		}
	}

}

// watchDirectory - watch the directory for configuration changes
func (p *Provider) watchDirectory(*chan ConfigurationUpdate) {

	// todo: implement directory watcher.

	// Watch for changes in the directory
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
	err = watcher.Add(p.Directory)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
