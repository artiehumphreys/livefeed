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
	mtx    sync.Mutex
	poller *Poller
	client *ingest.Client
}

func NewServer() *Server {
	return &Server{
		client: ingest.NewClient(),
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

	// replace active poller
	p := NewPoller(uint32(gameID64))
	p.Start(5 * time.Second)
	s.poller = p

	w.WriteHeader(http.StatusOK)
}

func (s *Server) SnapshotHandler(w http.ResponseWriter, r *http.Request) {
	// frontend polling
	s.mtx.Lock()
	p := s.poller
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
