package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

func (ra *RunAnalyzer) finalizeRun(end uint16) {
	if ra == nil {
		return
	}

	if ra.ActiveRun.PointsFor >= 7 {
		ra.Runs = append(ra.Runs, types.ScoringRun{
			TeamID:        ra.ActiveRun.TeamID,
			StartIndex:    ra.ActiveRun.StartIndex,
			EndIndex:      end,
			PointsFor:     ra.ActiveRun.PointsFor,
			PointsAgainst: ra.ActiveRun.PointsAgainst,
			IsKillShot:    ra.ActiveRun.PointsFor >= 10 && ra.ActiveRun.PointsAgainst == 0,
		})
	}

	ra.ActiveRun = nil
}

func (ra *RunAnalyzer) ProcessPlay(play *types.Play) {
	if ra.LastPlay == nil {
		ra.LastPlay = play
		return
	}

	deltaHome := play.HomeScore - ra.LastPlay.HomeScore
	deltaVisitor := play.VisitorScore - ra.LastPlay.VisitorScore

	if deltaHome == 0 && deltaVisitor == 0 {
		ra.LastPlay = play
		return
	}

	// malformed data
	if deltaHome > 0 && deltaVisitor > 0 {
		ra.LastPlay = play
		return
	}

	// BIG assumption: teamID corresponds to scoring team
	scoringTeam := play.TeamID
	if scoringTeam == 0 {
		ra.LastPlay = play
		return
	}

	var pointsScored = max(deltaHome, deltaVisitor)

	if ra.ActiveRun == nil {
		ra.ActiveRun = &types.ActiveScoringRun{
			TeamID:        scoringTeam,
			StartIndex:    play.Index,
			PointsFor:     pointsScored,
			PointsAgainst: 0,
		}
		ra.LastPlay = play
		return
	}

	if scoringTeam == ra.ActiveRun.TeamID {
		ra.ActiveRun.PointsFor += pointsScored
		ra.LastPlay = play
		return
	}

	// other team scored
	if ra.ActiveRun.PointsAgainst+pointsScored <= 2 {
		ra.ActiveRun.PointsAgainst += pointsScored
		ra.LastPlay = play
		return
	}

	ra.finalizeRun(play.Index - 1)
	ra.LastPlay = play
}

func ComputeRuns(pbp *types.PlayByPlaySummary) []types.ScoringRun {
	ra := &RunAnalyzer{}

	for i := range pbp.Plays {
		play := &pbp.Plays[i]
		play.Index = uint16(i)
		ra.ProcessPlay(play)
	}

	if ra.ActiveRun != nil {
		ra.finalizeRun(uint16(len(pbp.Plays) - 1))
	}

	return ra.Runs
}
