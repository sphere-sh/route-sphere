package route_sphere

type AddDomainConfigurationUpdater struct{}
type UpdateDomainConfigurationUpdater struct{}
type DeleteDomainConfigurationUpdater struct{}

// Update - Add a new domain configuration to the configurations.
func (a AddDomainConfigurationUpdater) Update(update ConfigurationUpdate) {
	// todo: implement update
}

// Update - Update an existing domain configuration in the configurations.
func (u UpdateDomainConfigurationUpdater) Update(update ConfigurationUpdate) {
	// todo: implement update
}

// Update - Delete an existing domain configuration from the configurations.
func (d DeleteDomainConfigurationUpdater) Update(update ConfigurationUpdate) {
	// todo: implement update
}
