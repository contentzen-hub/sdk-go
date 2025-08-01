// Package contentzen provides a Go SDK for ContentZen CMS.
// It supports both public and authenticated API endpoints.
package contentzen

import (
	"net/http"
	"time"
)

// Client is the main struct for interacting with ContentZen API.
type Client struct {
	BaseURL    string
	APIToken   string
	HTTPClient *http.Client
}

// NewClient creates a new ContentZen API client. If apiToken is empty, only public endpoints are available.
func NewClient(apiToken string) *Client {
	return &Client{
		BaseURL:    "https://api.contentzen.io",
		APIToken:   apiToken,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}
