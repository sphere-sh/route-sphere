package cli_utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// SessionGet - reads the session file and returns the session
func SessionGet() ([]string, error) {
	sessionFile, err := os.ReadFile("/etc/route-sphere/cli/session")
	if err != nil {
		return nil, fmt.Errorf("Failed to read session file: %w", err)
	}

	session := []string{}
	err = json.Unmarshal(sessionFile, &session)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal session file: %w", err)
	}

	return session, nil
}
