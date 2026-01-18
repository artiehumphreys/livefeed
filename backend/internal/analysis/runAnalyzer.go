package analysis

import "github.com/artiehumphreys/livefeed/internal/types"

type RunAnalyzer struct {
	LastPlay  *types.Play
	ActiveRun *types.ActiveScoringRun
	Runs      []types.ScoringRun
}
