package cli_utils

import "os"

// todo: work with environment variables instead of hardcoded paths

// GetSession - function returning the session string from filesystem.
func GetSession() (string, error) {
	content, err := os.ReadFile("/etc/route-sphere/cli/session")
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// HasSession - function to determine if session file exists.
func HasSession() bool {
	_, err := os.Stat("/etc/route-sphere/cli/session")
	return err == nil
}
