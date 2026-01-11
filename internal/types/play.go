package types

type RawPlay struct {
	TeamID           uint16 `json:"teamId"`
	HomeScore        uint16 `json:"homeScore"`
	AwayScore        uint16 `json:"awayScore"`
	Clock            string `json:"clock"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"LastName"`
	EventDescription string `json:"eventDescription"`
}
