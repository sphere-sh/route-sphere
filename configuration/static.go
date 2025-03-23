package configuration

type StaticConfiguration struct {
	Cloud struct {
		Enabled bool `json:"enabled"`
	} `json:"cloud"`
}

// CloudMode - returns the cloud mode status
func (s *StaticConfiguration) CloudMode() bool {
	return s.Cloud.Enabled
}
