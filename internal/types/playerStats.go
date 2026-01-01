package types

type RawPlayerStats struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Position  string `json:"position"`

	MP  string `json:"minutesPlayed"`
	FGM string `json:"fieldGoalsMade"`
	FGA string `json:"fieldGoalsAttempted"`
	FTM string `json:"freeThrowsMade"`
	FTA string `json:"freeThrowsAttempted"`
	TPM string `json:"threePointsMade"`
	TPA string `json:"threePointsAttempted"`

	OREB string `json:"offensiveRebounds"`
	REB  string `json:"totalRebounds"`
	AST  string `json:"assists"`
	TO   string `json:"turnovers"`
	PF   string `json:"personalFouls"`
	STL  string `json:"steals"`
	BLK  string `json:"blockedShots"`
	PTS  string `json:"points"`
}

type PlayerStats struct {
	ID        uint16
	Number    uint8
	FirstName string
	LastName  string
	Position  string

	MP  float32
	FGM uint8
	FGA uint8
	FTM uint8
	FTA uint8
	TPM uint8
	TPA uint8

	OREB uint8
	REB  uint8
	AST  uint8
	TO   uint8
	PF   uint8
	STL  uint8
	BLK  uint8
	PTS  uint8
}
