package types

import "encoding/json"

type RawPlay struct {
	TeamID           json.Number `json:"teamId"`
	IsHome           bool        `json:"isHome"`
	HomeScore        uint16      `json:"homeScore"`
	AwayScore        uint16      `json:"awayScore"`
	Clock            string      `json:"clock"`
	FirstName        string      `json:"firstName"`
	LastName         string      `json:"lastName"`
	EventDescription string      `json:"eventDescription"`
}

type Play struct {
	TeamID           uint16
	Period           uint8
	IsHome           bool
	HomeScore        uint16
	AwayScore        uint16
	Clock            *GameClock
	FirstName        string
	LastName         string
	EventDescription string
	IsScoringPlay    bool
}
