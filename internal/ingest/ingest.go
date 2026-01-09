package ingest

import (
	"log"

	"github.com/artiehumphreys/livefeed/internal/normalize"
)

const defaultGameID = "6502585"

func Run() {
	client := NewClient()

	data, err := client.FetchBoxScore(defaultGameID)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := normalize.ParseBoxScore(data)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := normalize.NormalizeBoxScore(raw)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", bs)
}
