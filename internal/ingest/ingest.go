package ingest

import (
	"log"

	"github.com/artiehumphreys/livefeed/internal/normalize"
)

const defaultGameID = 6502515

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

	pbpData, err := client.FetchPlayByPlay(defaultGameID)
	if err != nil {
		log.Fatal(err)
	}

	pbpRaw, err := normalize.ParsePlayByPlay(pbpData)
	if err != nil {
		log.Fatal(err)
	}

	pbp, err := normalize.NormalizePlayByPlay(pbpRaw)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v+\n", pbp)
}
