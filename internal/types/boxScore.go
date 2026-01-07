package types

type RawTeamBox struct {
	TeamID      string           `json:"teamId"`
	PlayerStats []RawPlayerStats `json:"playerStats"`
	TeamStats   RawTeamStats     `json:"teamStats"`
}

type RawBoxScore struct {
	ContestID  string       `json:"contestId"`
	Status     string       `json:"status"`
	Period     string       `json:"period"`
	Minutes    string       `json:"minutes"`
	Seconds    string       `json:"seconds"`
	SportsCode string       `json:"sportsCode"`
	Teams      []RawTeam    `json:"teams"`
	TeamBoxes  []RawTeamBox `json:"teamBoxscore"`
}

type BoxScore struct {
	ContestID int32
	Status    string
	// possibly null minutes and seconds
	Clock     *GameClock
	TeamBoxes []Team
}
