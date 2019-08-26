package endpass

import (
	"net/http"
	"time"
)

var (
	// Root endpoint for the public API
	PublicAPIBaseURL = "https://api.endpass.com/v1"

	// Root endpoint for the OAuth 2 API
	OAuth2BaseURL = "https://identity.endpass.com"
)

func defaultHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:   timeout,
	}
}
