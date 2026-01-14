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
	oppDREB := float32(opp.REB - opp.OREB)

	var orebPct float32
	if team.OREB+uint16(oppDREB) > 0 {
		orebPct = float32(team.OREB) / (float32(team.OREB) + oppDREB)
	}

	var drebPct float32
	if team.REB-team.OREB+opp.OREB > 0 {
		drebPct = float32(team.REB-team.OREB) /
			(float32(team.REB-team.OREB) + float32(opp.OREB))
	}

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
