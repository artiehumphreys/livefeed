package ingest

import (
	"fmt"
	"io"
	"net/http"
)

// https://pkg.go.dev/net/http#hdr-Clients_and_Transports
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// lol no oop
func NewClient() *Client {
	return &Client{
		BaseURL:    "https://ncaa-api.henrygd.me",
		HTTPClient: http.DefaultClient,
	}
}

// https://go.dev/blog/error-handling-and-go
func (c *Client) FetchBoxScore(gameId string) ([]byte, error) {
	url := fmt.Sprintf("%s/game/%s/boxscore", c.BaseURL, gameId)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to Fetch Data: GET request error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to Fetch Data: HTTP status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to Read Request Body.")
	}

	return body, nil
}
