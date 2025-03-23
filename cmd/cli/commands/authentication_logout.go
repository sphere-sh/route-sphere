package cli_commands

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type AuthenticationLogout struct {
}

func (cmd *AuthenticationLogout) Run(args *AuthenticationLogout, ctx *context.Context) {

	sessionFile, err := os.ReadFile("/etc/route-sphere/cli/session")
	if err != nil {
		slog.Error("Failed to read session file", "error", err)
		os.Exit(1)
	}

	session := []string{}
	err = json.Unmarshal(sessionFile, &session)
	if err != nil {
		slog.Error("Failed to parse session file", "error", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", "https://route.api.sphere.sh/logout", nil)
	if err != nil {
		slog.Error("Failed to create request", "error", err)
		os.Exit(1)
	}

	for _, cookie := range session {
		req.Header.Add("Cookie", cookie)
	}

	//body := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, args.Username, args.Password)
	//
	//client := &http.Client{}
	//req, err := http.NewRequest("POST", "https://route.api.sphere.sh/login", strings.NewReader(body))
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	//
	//// Add body to request
	////
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if resp.StatusCode != fiber.StatusNoContent {
	//	slog.Error("Failed to authenticate", "status", resp.Status)
	//	os.Exit(1)
	//}
	//
	//// Cookies to JSON.
	////
	//var cookies []string = make([]string, 0)
	//for _, cookie := range resp.Cookies() {
	//	cookies = append(cookies, cookie.String())
	//}
	//
	//// JSON encode cookies
	////
	//cookieJson, err := json.Marshal(cookies)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Write cookies to /etc/route-sphere/cli/session
	////
	//err = os.WriteFile("/etc/route-sphere/cli/session", cookieJson, 0644)
	//if err != nil {
	//	slog.Error("Failed to write session file", "error", err)
	//}
}
