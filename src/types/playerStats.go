package types

type RawPlayerStats struct {
	ID 	string `json:"id"`
	Number string `json:"number"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Position string `json:"position"`

	MP string `json:"minutesPlayed"`
	FGM string `json:"fieldGoalsMade"`
	FGA string `json:"fieldGoalsAttempted"`
	FTM string `json:"freeThrowsAttempted"`
	FTA string `json:"freeThrowsMade"`
	TPM string `json:"threePointsMade"`
	TPA string `json:"threePointsAttempted"`

	OREB string `json:"offensiveRebounds"`
	REB string `json:"totalRebounds"`
	AST string `json:"assists"`
	TO string `json:"turnovers"`
	PF string `json:"personalFouls"`
	ST string `json:"steals"`
	BLK string `json:"blockedShots"`
	PTS string `json:"points"`
}
