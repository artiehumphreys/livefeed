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
	teamID uint32,
) (types.RawTeamBox, bool) {
	for _, b := range boxes {
		if b.TeamID == teamID {
			return b, true
		}
	}
	return types.RawTeamBox{}, false
}

func getStatus(s string) types.GameStatus {
	switch s {
	case "F":
		return types.StatusFinal
	case "I":
		return types.StatusLive
	}

	return types.StatusBefore
}

func NormalizeBoxScore(raw *types.RawBoxScore) (*types.BoxScore, error) {
	clock := parseClockFromBoxScore(raw)

	teams := make([]types.Team, 0, len(raw.Teams))
	for _, rt := range raw.Teams {
		id := stou32(rt.TeamID)
		rb, ok := findTeamBox(raw.TeamBoxes, id)
		if !ok {
			return nil, fmt.Errorf("missing team box for teamId=%q", rt.TeamID)
		}

		teams = append(teams, parseTeam(rb, rt))
	}

	return &types.BoxScore{
		ContestID: raw.ContestID,
		Status:    getStatus(raw.Status),
		Period:    strings.TrimSpace(raw.Period),
		Clock:     clock,
		TeamBoxes: teams,
	}, nil
}
