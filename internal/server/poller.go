package server

import (
	"sync"
	"time"

	"github.com/artiehumphreys/livefeed/internal/analysis"
	"github.com/artiehumphreys/livefeed/internal/api"
	"github.com/artiehumphreys/livefeed/internal/ingest"
)

type Poller struct {
	gameID   uint32
	client   *ingest.Client
	mtx      sync.RWMutex
	snapshot *api.GameSnapshot
}

func NewPoller(gameID uint32) *Poller {
	return &Poller{
		gameID: gameID,
		client: ingest.NewClient(),
	}
}

func (p *Poller) PollOnce() {
	box, err := p.client.GetBoxScore(p.gameID)
	if err != nil {
		return
	}

	pbp, err := p.client.GetPlayByPlay(p.gameID)
	if err != nil {
		return
	}

	teams := make([]api.TeamSnapshot, 0, len(box.TeamBoxes))

	if len(box.TeamBoxes) == 2 {
		teamA := box.TeamBoxes[0]
		teamB := box.TeamBoxes[1]

		teams = append(teams, api.TeamSnapshot{
			TeamID:  teamA.TeamID,
			Name:    teamA.Name,
			Metrics: analysis.ComputeTeamMetrics(teamA.TeamID, teamA.Stats, teamB.Stats),
		}, api.TeamSnapshot{
			TeamID:  teamB.TeamID,
			Name:    teamB.Name,
			Metrics: analysis.ComputeTeamMetrics(teamB.TeamID, teamB.Stats, teamA.Stats),
		})
	}

	runs := analysis.ComputeRuns(pbp)

	// atomic snapshots
	p.mtx.Lock()
	p.snapshot = &api.GameSnapshot{
		ContestID:   p.gameID,
		Teams:       teams,
		BoxScore:    box,
		PlayByPlay:  pbp,
		Runs:        runs,
		LastUpdated: time.Now().Unix(),
	}
	p.mtx.Unlock()
}
