package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

type TeamMetrics struct {
	TeamID uint16

	PPP         float32
	TurnoverPct float32

	OffensiveReboundPct float32
	DefensiveReboundPct float32

	ThreePointAttemptRate float32
}

const Y float32 = 0.44

func getPossessions(team types.TeamStats) float32 {
	oreb := float32(team.OREB)
	fta := float32(team.FTA)
	to := float32(team.TO)
	fga := float32(team.FGA)

	possessions := fga - oreb + to + Y*fta
	if possessions <= 0 {
		return 0
	}

	return possessions
}

func PointsPerPossession(team types.TeamStats) float32 {
	// https://www.hoopcoach.org/points-per-possessions/
	possessions := getPossessions(team)
	if possessions <= 0 {
		return 0
	}

	return float32(team.PTS) / possessions
}

func TurnoverPercentage(team types.TeamStats) float32 {
	possessions := getPossessions(team)
	if possessions <= 0 {
		return 0
	}

	return float32(team.TO) / possessions
}

func OffensiveReboundPercentage(team, opp types.TeamStats) float32 {
	oppDREB := float32(opp.REB) - float32(opp.OREB)
	if team.OREB == 0 && oppDREB == 0 {
		return 0
	}
	return float32(team.OREB) / (float32(team.OREB) + oppDREB)
}

func DefensiveReboundPercentage(team, opp types.TeamStats) float32 {
	teamDREB := float32(team.REB) - float32(team.OREB)
	oppOREB := float32(opp.OREB)
	if teamDREB+oppOREB == 0 {
		return 0
	}
	return teamDREB / (teamDREB + oppOREB)
}

func ThreePointAttemptRate(team types.TeamStats) float32 {
	if team.FGA == 0 {
		return 0
	}
	return float32(team.TPA) / float32(team.FGA)
}

func ComputeTeamMetrics(teamID uint16, team, opp types.TeamStats) TeamMetrics {
	return TeamMetrics{
		TeamID:                teamID,
		PPP:                   PointsPerPossession(team),
		TurnoverPct:           TurnoverPercentage(team),
		OffensiveReboundPct:   OffensiveReboundPercentage(team, opp),
		DefensiveReboundPct:   DefensiveReboundPercentage(team, opp),
		ThreePointAttemptRate: ThreePointAttemptRate(team),
	}
}
