package cli_utils

import "os"

// ApiGetBaseUrl - function used to get the base URL for the API.
func ApiGetBaseUrl() string {
	envUrl := os.Getenv("ROUTE_SPHERE_API_URL")
	if envUrl == "" {
		return "https://route.api.sphere.sh"
	}
	return envUrl
}
