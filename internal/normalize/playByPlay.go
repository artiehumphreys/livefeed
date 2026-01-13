package normalize

import (
	"encoding/json"
	"fmt"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func ParsePlayByPlay(data []byte) (*types.RawPlayByPlaySummary, error) {
	var res types.RawPlayByPlaySummary

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func NormalizePlayByPlay(raw *types.RawPlayByPlaySummary) (*types.PlayByPlaySummary, error) {
	periods, err := NormalizePeriods(raw.Periods)
	if err != nil {
		return nil, fmt.Errorf("error parsing periods %d: %w", raw.ContestID, err)
	}

	summary := &types.PlayByPlaySummary{
		Plays:     make([]types.Play, 0),
		HomePlays: make([]uint16, 0),
		AwayPlays: make([]uint16, 0),
	}

	for _, period := range periods {
		for _, play := range period.PlayByPlay {
			// could append entire slice to avoid multiple reallocations
			summary.Plays = append(summary.Plays, play)

			idx := uint16(len(summary.Plays))

			if play.IsHome {
				summary.HomePlays = append(summary.HomePlays, idx)
			} else {
				summary.AwayPlays = append(summary.AwayPlays, idx)
			}
		}
	}

	return summary, nil
}
