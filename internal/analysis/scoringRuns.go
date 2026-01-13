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

func (ra *RunAnalyzer) ProcessPlay(play types.Play) {
	if ra.LastPlay == nil {
		ra.LastPlay = &play
		return
	}

	deltaHome := play.HomeScore - ra.LastPlay.HomeScore
	deltaVisitor := play.VisitorScore - ra.LastPlay.VisitorScore
}
