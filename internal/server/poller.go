package server

import (
	"sync"

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
