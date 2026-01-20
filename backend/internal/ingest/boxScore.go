package ingest

import (
	"fmt"
	"io"
	"net/http"
)

// https://go.dev/blog/error-handling-and-go
func (c *Client) FetchBoxScore(gameId uint32) ([]byte, error) {
	url := fmt.Sprintf("%s/game/%d/boxscore", c.BaseURL, gameId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "livefeed/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body from %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		snippet := string(body)
		if len(snippet) > 300 {
			snippet = snippet[:300]
		}
		return nil, fmt.Errorf("GET %s -> %d, body: %q", url, resp.StatusCode, snippet)
	}

	return body, nil
}
