package endpass

import "golang.org/x/oauth2"

var (
	// Root endpoint for the API
	APIBase = "https://api.endpass.com/v1"

	OAuth2BaseURL = "https://identity-dev.endpass.com"

	OAuth2Endpoint = oauth2.Endpoint{
		AuthURL:  OAuth2BaseURL + "/api/v1.1/oauth/auth",
		TokenURL: OAuth2BaseURL + "/api/v1.1/oauth/token",
	}
)
