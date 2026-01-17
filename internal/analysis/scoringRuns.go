package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

func clockToSeconds(c *types.GameClock) uint16 {
	if c == nil {
		return 0
	}
	return uint16(c.Minutes)*60 + uint16(c.Seconds)
}

func ComputeClockDiff(start *types.GameClock, end *types.GameClock, startPeriod uint8, endPeriod uint8) uint16 {
	if start == nil || end == nil {
		return 0
	}

	startSec := clockToSeconds(start)
	endSec := clockToSeconds(end)

	if startPeriod == endPeriod {
		if startSec >= endSec {
			return startSec - endSec
		}
		return 0
	}

	const periodLength = 20 * 60

	diff := uint16(startSec)

	if endPeriod > startPeriod+1 {
		fullPeriods := endPeriod - startPeriod - 1
		diff += uint16(fullPeriods) * periodLength
	}

	if endSec <= periodLength {
		diff += periodLength - endSec
	}

	return diff
}

func (ra *RunAnalyzer) finalizeRun(end uint16, endClock *types.GameClock, endPeriod uint8) {
	if ra == nil {
		return
	}

	duration := ComputeClockDiff(
		ra.ActiveRun.StartClock,
		endClock,
		ra.ActiveRun.StartPeriod,
		endPeriod,
	)

	if ra.ActiveRun.PointsFor >= 7 {
		ra.Runs = append(ra.Runs, types.ScoringRun{
			TeamID:          ra.ActiveRun.TeamID,
			StartIndex:      ra.ActiveRun.StartIndex,
			EndIndex:        end,
			PointsFor:       ra.ActiveRun.PointsFor,
			PointsAgainst:   ra.ActiveRun.PointsAgainst,
			IsKillShot:      ra.ActiveRun.PointsFor >= 10 && ra.ActiveRun.PointsAgainst == 0,
			DurationSeconds: duration,
			Period:          ra.ActiveRun.StartPeriod,
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

	play.IsScoringPlay = true

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
			StartPeriod:   play.Period,
			StartClock:    play.Clock,
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

	ra.finalizeRun(play.Index-1, play.Clock, play.Period)
	ra.LastPlay = play
}

func ComputeRuns(pbp *types.PlayByPlaySummary) []types.ScoringRun {
	ra := &RunAnalyzer{}

	for i := range pbp.Plays {
		play := &pbp.Plays[i]
		play.Index = uint16(i)
		ra.ProcessPlay(play)
	}

	if ra.ActiveRun != nil && len(pbp.Plays) > 0 {
		last := &pbp.Plays[len(pbp.Plays)-1]
		ra.finalizeRun(last.Index, last.Clock, last.Period)
	}

	return ra.Runs
}
