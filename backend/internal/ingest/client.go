package ingest

import (
	"net/http"
	"time"
)

// https://pkg.go.dev/net/http#hdr-Clients_and_Transports
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// lol no oop
func NewClient() *Client {
	return &Client{
		BaseURL: "https://ncaa-api.henrygd.me",
		HTTPClient: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}
