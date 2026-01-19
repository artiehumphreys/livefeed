package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/artiehumphreys/livefeed/internal/api"
	"github.com/artiehumphreys/livefeed/internal/ingest"
)

type ScoreboardPoller struct {
	client   *ingest.Client
	mtx      sync.RWMutex
	snapshot *api.ScoreboardSnapshot
}

func NewScoreboardPoller() *ScoreboardPoller {
	return &ScoreboardPoller{
		client: ingest.NewClient(),
	}
}

func (s *ScoreboardPoller) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		// immediately run once
		s.PollOnce()

		for range ticker.C {
			s.PollOnce()
		}
	}()
}

func (s *ScoreboardPoller) PollOnce() {
	fmt.Println("Polling scoreboard")

	games, err := s.client.GetScoreboard()
	if err != nil {
		fmt.Println("Error fetching scoreboard:", err)
		return
	}

	s.mtx.Lock()
	s.snapshot = &api.ScoreboardSnapshot{
		Games:       games,
		LastUpdated: time.Now().Unix(),
	}
	s.mtx.Unlock()

	fmt.Println("Scoreboard updated")
}

func (s *ScoreboardPoller) GetSnapshot() *api.ScoreboardSnapshot {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.snapshot
}
