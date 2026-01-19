package normalize

import (
	"encoding/json"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func ParseScoreboard(data []byte) (*types.RawScoreboardResponse, error) {
	var res types.RawScoreboardResponse
	// short initialization to reduce scope of error, neat pattern
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func NormalizeScoreboard(raw []types.RawScoreboardGames) []types.GameSummary {
	res := make([]types.GameSummary, 0, len(raw))

	for _, g := range raw {
		game := g.Game

		gameID := stou32(game.GameID)

		homeScore := stou16(strings.TrimSpace(game.Home.Score))
		awayScore := stou16(strings.TrimSpace(game.Away.Score))

		res = append(res, types.GameSummary{
			GameID:   gameID,
			HomeTeam: game.Home.Name.Short,
			AwayTeam: game.Away.Name.Short,

			HomeScore: homeScore,
			AwayScore: awayScore,

			State: game.GameState,

			StartTime: strings.TrimSpace(game.StartTime),
			Clock:     strings.TrimSpace(game.ContestClock),
		})
	}

	return res
}
