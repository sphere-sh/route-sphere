package cloud

import (
	"encoding/json"
	"net/http"
	"os"
)

func init() {
	os.MkdirAll("/etc/route-sphere/cloud", 0755)
}

type Session struct {
	Cookies []string
}

// SessionGet - returns initialized session.
func SessionGet() *Session {
	return &Session{
		Cookies: getCookiesFromStorage(),
	}
}

// getCookiesFromStorage - returns cookies from filesytem.
func getCookiesFromStorage() []string {
	sessionFile, err := os.ReadFile("/etc/route-sphere/cloud/session")
	if err != nil {
		return []string{}
	}

	session := []string{}
	err = json.Unmarshal(sessionFile, &session)
	if err != nil {
		return []string{}
	}

	return session
}

// Cookies2Request - injects cookies into request.
func (s *Session) Cookies2Request(request *http.Request) {
	for _, cookie := range s.Cookies {
		request.Header.Add("Cookie", cookie)
	}
}
