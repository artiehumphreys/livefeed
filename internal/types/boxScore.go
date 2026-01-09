package types

import "encoding/json"

type RawTeamBox struct {
	TeamID      string           `json:"teamId"`
	PlayerStats []RawPlayerStats `json:"playerStats"`
	TeamStats   RawTeamStats     `json:"teamStats"`
}

type RawBoxScore struct {
	ContestID  json.Number  `json:"contestId"`
	Status     string       `json:"status"`
	Period     string       `json:"period"`
	Minutes    *json.Number `json:"minutes"`
	Seconds    *json.Number `json:"seconds"`
	SportsCode string       `json:"sportsCode"`
	Teams      []RawTeam    `json:"teams"`
	TeamBoxes  []RawTeamBox `json:"teamBoxscore"`
}

type BoxScore struct {
	ContestID uint32
	Status    string
	Period    string
	// possibly null minutes and seconds
	Clock     *GameClock
	TeamBoxes []Team
}
