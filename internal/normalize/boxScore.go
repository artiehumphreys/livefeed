package normalize

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func ParseBoxScore(data []byte) (*types.RawBoxScore, error) {
	var res types.RawBoxScore
	// short initialization to reduce scope of error, neat pattern
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func findTeamBox(
	boxes []types.RawTeamBox,
	teamID string,
) (types.RawTeamBox, bool) {
	for _, b := range boxes {
		if b.TeamID == teamID {
			return b, true
		}
	}
	return types.RawTeamBox{}, false
}

func NormalizeBoxScore(raw *types.RawBoxScore) (*types.BoxScore, error) {
	contestID, err := raw.ContestID.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid contestId %q: %w", raw.ContestID, err)
	}

	clock := parseClock(raw)

	teams := make([]types.Team, 0, len(raw.Teams))
	for _, rt := range raw.Teams {
		rb, ok := findTeamBox(raw.TeamBoxes, rt.TeamID)
		if !ok {
			return nil, fmt.Errorf("missing team box for teamId=%s", rt.TeamID)
		}

		teams = append(teams, parseTeam(rb, rt))
	}

	return &types.BoxScore{
		ContestID: uint32(contestID),
		Status:    strings.TrimSpace(raw.Status),
		Period:    strings.TrimSpace(raw.Period),
		Clock:     clock,
		TeamBoxes: teams,
	}, nil
}
