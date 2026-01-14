package ingest

import (
	"log"

	"github.com/artiehumphreys/livefeed/internal/normalize"
	"github.com/artiehumphreys/livefeed/internal/types"
)

func (c *Client) GetBoxScore(gameID uint32) *types.BoxScore {
	data, err := c.FetchBoxScore(gameID)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := normalize.ParseBoxScore(data)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := normalize.NormalizeBoxScore(raw)
	if err != nil {
		log.Fatal(err)
	}

	return bs
}

func (c *Client) GetPlayByPlay(gameID uint32) *types.PlayByPlaySummary {
	data, err := c.FetchPlayByPlay(gameID)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := normalize.ParsePlayByPlay(data)
	if err != nil {
		log.Fatal(err)
	}

	pbp, err := normalize.NormalizePlayByPlay(raw)
	if err != nil {
		log.Fatal(err)
	}

	return pbp
}
