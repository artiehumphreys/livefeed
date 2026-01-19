package types

type RawScoreboardResponse struct {
	Games []RawScoreboardGames `json:"games"`
}

type RawScoreboardGames struct {
	Game RawGame `json:"game"`
}

type RawGame struct {
	GameID        string `json:"gameID"`
	GameState     string `json:"gameState"`
	StartTime     string `json:"startTime"`
	CurrentPeriod string `json:"currentPeriod"`
	ContestClock  string `json:"contestClock"`

	Home RawSide `json:"home"`
	Away RawSide `json:"away"`
}

type RawSide struct {
	Score string `json:"score"`
	Name  struct {
		Short string `json:"short"`
	} `json:"names"`
}

type GameSummary struct {
	GameID   uint32
	HomeTeam string
	AwayTeam string

	HomeScore uint16
	AwayScore uint16

	State string // pre, live, or final

	StartTime string
	Clock     string
}
