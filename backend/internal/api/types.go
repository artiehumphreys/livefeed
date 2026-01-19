package api

import (
	"github.com/artiehumphreys/livefeed/internal/analysis"
	"github.com/artiehumphreys/livefeed/internal/types"
)

type GameSnapshot struct {
	ContestID   uint32
	Teams       []TeamSnapshot
	BoxScore    *types.BoxScore
	PlayByPlay  *types.PlayByPlaySummary
	Runs        []types.ScoringRun
	LastUpdated int64
}

type TeamSnapshot struct {
	TeamID  uint16
	Name    string
	Metrics analysis.TeamMetrics
}

type ScoreboardSnapshot struct {
	Games       []types.GameSummary
	LastUpdated int64
}
