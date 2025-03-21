package cli_commands

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	cli_utils "route-sphere/cmd/cli/utils"
)

type ConnectionInstall struct {
	Id             string `arg:"positional,required"`
	PrivateKeyPath string `arg:"required,-p,--private-key" help:"path to server private key"`
}

func (cmd *ConnectionInstall) Run(args *ConnectionInstall) {

	session, err := cli_utils.SessionGet()
	if err != nil {
		println("Failed to get session")
		return
	}

	// Get contents of the private key file
	//
	privateKey, err := os.ReadFile(args.PrivateKeyPath)
	if err != nil {
		println("Failed to read private key file")
		return
	}

	// Write private key to the server
	//
	err = os.WriteFile("/etc/route-sphere/server/tls/server.key", privateKey, 0644)
	if err != nil {
		println("Failed to write private key to file")
		return
	}

	downloadIntermediateCA(session, args)
	downloadServerCertificate(session, args)
}

// downloadServerCertificate - downloads the server certificate from the API
func downloadServerCertificate(session []string, args *ConnectionInstall) {
	// Download server certificate from the API
	//
	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/api/v1/connection/"+args.Id+"/server-certificate/download", nil)
	if err != nil {
		println("Failed to create request")
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for _, cookie := range session {
		req.Header.Add("Cookie", cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to get connections", "error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}

	// Create `server/tls` directory
	//
	err = os.MkdirAll("/etc/route-sphere/server/tls", 0755)
	if err != nil {
		println("Failed to create directory")
		return
	}

	// Store server certificate in a file
	//
	err = os.WriteFile("/etc/route-sphere/server/tls/server.crt", bodyBytes, 0644)
	if err != nil {
		println("Failed to write server certificate to file")
		return
	}
}

// downloadIntermediateCA - downloads the intermediate CA (chain) from the API
func downloadIntermediateCA(session []string, args *ConnectionInstall) {
	// Download intermediate CA (chain) from the API
	//
	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/api/v1/connection/"+args.Id+"/intermediate-ca/download", nil)
	if err != nil {
		println("Failed to create request")
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for _, cookie := range session {
		req.Header.Add("Cookie", cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to get connections", "error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		os.Exit(1)
	}

	// Create `server/tls` directory
	//
	err = os.MkdirAll("/etc/route-sphere/server/tls", 0755)
	if err != nil {
		println("Failed to create directory")
		return
	}

	// Store intermediate CA (chain) in a file
	//
	err = os.WriteFile("/etc/route-sphere/server/tls/intermediate-ca.pem", bodyBytes, 0644)
	if err != nil {
		println("Failed to write intermediate CA to file")
		return
	}
}
