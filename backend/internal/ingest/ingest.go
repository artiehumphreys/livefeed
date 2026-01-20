package ingest

import (
	"fmt"

	"github.com/artiehumphreys/livefeed/internal/normalize"
	"github.com/artiehumphreys/livefeed/internal/types"
)

func (c *Client) GetBoxScore(gameID uint32) (*types.BoxScore, error) {
	data, err := c.FetchBoxScore(gameID)
	if err != nil {
		fmt.Println("unable to read box score")
		return nil, err
	}

	raw, err := normalize.ParseBoxScore(data)
	if err != nil {
		return nil, err
	}

	bs, err := normalize.NormalizeBoxScore(raw)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
func (c *Client) GetPlayByPlay(gameID uint32) (*types.PlayByPlaySummary, error) {
	data, err := c.FetchPlayByPlay(gameID)
	if err != nil {
		fmt.Println("unable to read play by play")
		return nil, err
	}

	raw, err := normalize.ParsePlayByPlay(data)
	if err != nil {
		return nil, err
	}

	pbp, err := normalize.NormalizePlayByPlay(raw)
	if err != nil {
		return nil, err
	}

	return pbp, nil
}

func (c *Client) GetScoreboard() ([]types.GameSummary, error) {
	data, err := c.FetchScoreboard()
	if err != nil {
		fmt.Println("unable to read scoreboard")
		return nil, err
	}

	raw, err := normalize.ParseScoreboard(data)
	if err != nil {
		return nil, err
	}

	sb, err := normalize.NormalizeScoreboard(raw.Games)
	if err != nil {
		return nil, err
	}

	return sb, nil
}
