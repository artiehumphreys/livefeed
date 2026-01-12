package types

type RawPlayByPlaySummary struct {
	ContestID uint32      `json:"contestId"`
	Status    string      `json:"status"`
	Periods   []RawPeriod `json:"periods"`
}

type PlayByPlaySummary struct {
	Plays []Play
	// storing indices of plays
	HomePlays []uint16
	AwayPlays []uint16
}
