package server

import (
	"fmt"
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

	ticker *time.Ticker
	stop   chan struct{}
}

func (p *Poller) Start(interval time.Duration) {
	if p.stop != nil {
		return
	}
	p.stop = make(chan struct{})
	p.ticker = time.NewTicker(interval)

	// run a new goroutine to fetch information every `interval` seconds
	go func() {
		p.PollOnce()
		for {
			select {
			case <-p.ticker.C:
				p.PollOnce()
			case <-p.stop:
				return
			}
		}
	}()
}

func (p *Poller) Stop() {
	if p.stop == nil {
		return
	}
	close(p.stop)
	p.stop = nil
	if p.ticker != nil {
		p.ticker.Stop()
		p.ticker = nil
	}
}

func NewPoller(gameID uint32) *Poller {
	return &Poller{
		gameID: gameID,
		client: ingest.NewClient(),
	}
}

func (p *Poller) GetSnapshot() *api.GameSnapshot {
	// ensure no race conditions with data
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.snapshot
}

func (p *Poller) PollOnce() {
	fmt.Println("Polling game", p.gameID)

	box, err := p.client.GetBoxScore(p.gameID)
	if err != nil {
		fmt.Println("Error fetching box score:", err)
		return
	}

	pbp, err := p.client.GetPlayByPlay(p.gameID)
	if err != nil {
		fmt.Println("Error fetching play-by-play:", err)
		return
	}

	fmt.Println("Data fetched successfully!")
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
	defer p.mtx.Unlock()

	p.snapshot = &api.GameSnapshot{
		ContestID:   p.gameID,
		Teams:       teams,
		BoxScore:    box,
		PlayByPlay:  pbp,
		Runs:        runs,
		LastUpdated: time.Now().Unix(),
	}

	fmt.Println("Snapshot updated!")
}
