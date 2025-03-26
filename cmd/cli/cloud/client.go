package cloud

import (
	"encoding/json"
	"net/http"
	"os"
)

type CloudHTTPClient struct {
	// Cookies - HTTP cookies containing the session information
	Cookies []*http.Cookie

	// Client - HTTP client to make requests
	Client *http.Client

	// BaseURL - Base URL for the cloud
	BaseURL string
}

// NewCloudHTTPClient - Create a new HTTP client for the cloud
func NewCloudHTTPClient() *CloudHTTPClient {

	// Read session file from disk
	//
	var cookies = []*http.Cookie{}
	session, err := os.ReadFile("/etc/route-sphere/cli/session")
	if err != nil {
		panic(err)
	}

	// Unmarshal session file
	//
	err = json.Unmarshal(session, &cookies)
	if err != nil {
		panic(err)
	}

	return &CloudHTTPClient{
		Client:  &http.Client{},
		BaseURL: os.Getenv("ROUTE_SPHERE_API_BASE_URL"),
		Cookies: cookies,
	}
}

// Get - Make a GET request to the cloud
func (c *CloudHTTPClient) Get(url string) (*http.Response, error) {
	println(c.BaseURL + url)
	req, err := http.NewRequest("GET", c.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add cookies
	//
	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	println(resp.Status)

	return resp, nil
}
