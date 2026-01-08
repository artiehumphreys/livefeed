package normalize

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseBoxScore(data []byte) (*types.RawBoxScore, error) {
	var res types.RawBoxScore
	// short initialization to reduce scope of error, neat pattern
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func normalizeBoxScore(raw *types.RawBoxScore) (*types.BoxScore, error) {
	contestStr := strings.TrimSpace(raw.ContestID)
	contestID, err := strconv.ParseInt(contestStr, 10, 32)

	if err != nil {
		return nil, fmt.Errorf("invalid contest ID %q: %w", contestStr, err)
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
