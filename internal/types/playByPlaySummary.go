package types

type PlayByPlaySummary struct {
	Plays []Play
	// storing indices of plays
	HomePlays []uint16
	AwayPlays []uint16
}
