package skroutz

import "net/http"

// Config struct for the Skroutz client.
type Config struct {
	HTTPClient  *http.Client
	AccessToken string
}

// NewConfig requires accessToken.
func NewConfig(accessToken string) *Config {
	return &Config{
		HTTPClient:  http.DefaultClient,
		AccessToken: accessToken,
	}
}
