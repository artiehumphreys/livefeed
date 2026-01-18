package main

import (
	"log"
	"net/http"

	"github.com/artiehumphreys/livefeed/internal/server"
)

func main() {
	srv := server.NewServer()

	http.HandleFunc("/set-game", srv.SetGameHandler)
	http.HandleFunc("/snapshot", srv.SnapshotHandler)
	http.HandleFunc("/validate-game", srv.ValidateGameHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
