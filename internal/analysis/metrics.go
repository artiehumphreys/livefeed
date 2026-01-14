package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

type TeamMetrics struct {
	PPP float32

	OffensiveReboundPct float32
	DefensiveReboundPct float32

	ThreePointAttemptRate float32
}

func PointsPerPossession(team types.TeamStats) float32 {
	// https://www.hoopcoach.org/points-per-possessions/
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

func ComputeTeamMetrics(team, opp types.TeamStats) TeamMetrics {
	// Offensive Rebound %
	oppDREB := float32(opp.REB) - float32(opp.OREB)

	var orebPct float32
	if team.OREB > 0 || oppDREB > 0 {
		orebPct = float32(team.OREB) / (float32(team.OREB) + oppDREB)
	}

	// Defensive Rebound %
	teamDREB := float32(team.REB) - float32(team.OREB)
	oppOREB := float32(opp.OREB)

	var drebPct float32
	if teamDREB+oppOREB > 0 {
		drebPct = teamDREB / (teamDREB + oppOREB)
	}

	// 3PA rate
	var threeAttemptRate float32
	if team.FGA > 0 {
		threeAttemptRate = float32(team.TPA) / float32(team.FGA)
	}

	return TeamMetrics{
		PPP:                   PointsPerPossession(team),
		OffensiveReboundPct:   orebPct,
		DefensiveReboundPct:   drebPct,
		ThreePointAttemptRate: threeAttemptRate,
	}
}
