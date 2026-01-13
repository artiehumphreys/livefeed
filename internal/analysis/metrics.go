package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

func PointsPerPossession(team types.TeamStats) float32 {
	possessions :=
		float32(team.FGA) +
			float32(team.TO) +
			0.44*float32(team.FTA) -
			float32(team.OREB)

	if possessions <= 0 {
		return 0
	}

	return float32(team.PTS) / possessions
}
