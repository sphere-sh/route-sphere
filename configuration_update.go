package route_sphere

import (
	"github.com/sphere-sh/go-struct-sync/compare"
	"log/slog"
	"sync"
)

type ConfigurationUpdate struct {
	// Action - action to perform on the configuration
	Action string

	// Key - key of the configuration to execute the action on.
	Key string

	// Value - value of the configuration update (optional)
	Value interface{}
}

// InMemoryConfigurationUpdater - interface for struct that updates in-memory configuration.
type InMemoryConfigurationUpdater interface {
	// Update - update the in-memory configuration
	Update(update ConfigurationUpdate)
}

// ListenForConfigurationUpdates - listen for configuration updates
func (c *Configuration) ListenForConfigurationUpdates(wg *sync.WaitGroup) {
	defer wg.Done()

	updateImplementations := make(map[string]InMemoryConfigurationUpdater)
	updateImplementations["domains.ADD"] = AddDomainConfigurationUpdater{}

	// Compare the in-memory configuration with the provided configuration.
	res, err := compare.CompareStructs(c, Configuration{})
	if err != nil {
		slog.Error("Error comparing configuration", "error", err)
		return
	}

	slog.Info("Configuration updates found", "result", res)

	for {
		select {
		case update := <-c.UpdateChannel:
			key := update.Key + "." + update.Action
			if updater, ok := updateImplementations[key]; ok {
				updater.Update(update)
				continue
			}
			slog.Error("Unknown configuration update", "key", key)
		}
	}
}

// compare - compares the in-memory configuration with the provided configuration.
func (c *Configuration) compare() {

}
