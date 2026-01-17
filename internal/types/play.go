package types

import "encoding/json"

type RawPlay struct {
	TeamID           json.Number `json:"teamId"`
	IsHome           bool        `json:"isHome"`
	HomeScore        uint16      `json:"homeScore"`
	VisitorScore     uint16      `json:"visitorScore"`
	Clock            string      `json:"clock"`
	FirstName        string      `json:"firstName"`
	LastName         string      `json:"lastName"`
	EventDescription string      `json:"eventDescription"`
}

type Play struct {
	Index            uint16
	TeamID           uint16
	Period           uint8
	IsHome           bool
	HomeScore        uint16
	VisitorScore     uint16
	Clock            *GameClock
	FirstName        string
	LastName         string
	EventDescription string
	IsScoringPlay    bool
}
