package cli_commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type AuthenticationLogin struct {
	Username string `arg:"positional,required"`
	Password string `arg:"positional,required"`
}

func (cmd *AuthenticationLogin) Run(args *AuthenticationLogin, ctx *context.Context) {
	body := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, args.Username, args.Password)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://route.api.sphere.sh/login", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add body to request
	//

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != fiber.StatusNoContent {
		slog.Error("Failed to authenticate", "status", resp.Status)
		os.Exit(1)
	}

	// Cookies to JSON.
	//
	var cookies []string = make([]string, 0)
	for _, cookie := range resp.Cookies() {
		cookies = append(cookies, cookie.String())
	}

	// JSON encode cookies
	//
	cookieJson, err := json.Marshal(cookies)
	if err != nil {
		panic(err)
	}

	// Write cookies to /etc/route-sphere/cli/session
	//
	err = os.WriteFile("/etc/route-sphere/cli/session", cookieJson, 0644)
	if err != nil {
		slog.Error("Failed to write session file", "error", err)
	}

	os.MkdirAll("/etc/route-sphere/cloud", 0755)

	err = os.WriteFile("/etc/route-sphere/cloud/session", cookieJson, 0644)
	if err != nil {
		slog.Error("Failed to write session file", "error", err)
	}

	slog.Info("Successfully authenticated")
}
