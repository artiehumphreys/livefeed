package ingest

import (
	"log"

	"github.com/artiehumphreys/livefeed/internal/normalize"
	"github.com/artiehumphreys/livefeed/internal/types"
)

func (c *Client) GetBoxScore(gameID uint32) (*types.BoxScore, error) {
	data, err := c.FetchBoxScore(gameID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	raw, err := normalize.ParseBoxScore(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	bs, err := normalize.NormalizeBoxScore(raw)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return bs, nil
}

func (c *Client) GetPlayByPlay(gameID uint32) (*types.PlayByPlaySummary, error) {
	data, err := c.FetchPlayByPlay(gameID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	raw, err := normalize.ParsePlayByPlay(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pbp, err := normalize.NormalizePlayByPlay(raw)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return pbp, nil
}
