package main

import (
	"log"
	"net/http"

	"github.com/artiehumphreys/livefeed/internal/server"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func main() {
	srv := server.NewServer()

	http.HandleFunc("/scoreboard", withCORS(srv.ScoreboardHandler))
	http.HandleFunc("/snapshot", withCORS(srv.SnapshotHandler))
	http.HandleFunc("/set-game", withCORS(srv.SetGameHandler))
	http.HandleFunc("/validate-game", withCORS(srv.ValidateGameHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
