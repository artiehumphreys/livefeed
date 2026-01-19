package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/artiehumphreys/livefeed/internal/ingest"
)

type Server struct {
	mtx              sync.Mutex
	gamePoller       *Poller
	scoreboardPoller *ScoreboardPoller
	client           *ingest.Client
}

func NewServer() *Server {
	sb := NewScoreboardPoller()
	sb.Start(1 * time.Minute)

	return &Server{
		client:           ingest.NewClient(),
		scoreboardPoller: sb,
	}
}

func (s *Server) SetGameHandler(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.URL.Query().Get("gameId")
	if gameIDStr == "" {
		http.Error(w, "missing gameId", http.StatusBadRequest)
		return
	}

	gameID64, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid gameId", http.StatusBadRequest)
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// replace active gamePoller
	p := NewPoller(uint32(gameID64))
	p.Start(5 * time.Second)
	s.gamePoller = p

	w.WriteHeader(http.StatusOK)
}

func (s *Server) SnapshotHandler(w http.ResponseWriter, r *http.Request) {
	// frontend polling
	s.mtx.Lock()
	p := s.gamePoller
	s.mtx.Unlock()

	if p == nil {
		http.Error(w, "no active game", http.StatusServiceUnavailable)
		return
	}

	snap := p.GetSnapshot()
	if snap == nil {
		http.Error(w, "snapshot not ready", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snap)
}

func (s *Server) ValidateGameHandler(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.URL.Query().Get("gameId")
	if gameIDStr == "" {
		http.Error(w, "missing gameId", http.StatusBadRequest)
		return
	}

	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid gameId", http.StatusBadRequest)
		return
	}

	box, err := s.client.GetBoxScore(uint32(gameID))
	if err != nil || box == nil {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	}

	resp := map[string]any{
		"contestId": gameID,
		"status":    box.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) ScoreboardHandler(w http.ResponseWriter, r *http.Request) {
	snapshot := s.scoreboardPoller.GetSnapshot()
	if snapshot == nil {
		http.Error(w, "scoreboard not ready", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshot)
}
