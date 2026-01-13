package ingest

import (
	"fmt"
	"io"
	"net/http"
)

// https://go.dev/blog/error-handling-and-go
func (c *Client) FetchBoxScore(gameId uint32) ([]byte, error) {
	url := fmt.Sprintf("%s/game/%d/boxscore", c.BaseURL, gameId)

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
