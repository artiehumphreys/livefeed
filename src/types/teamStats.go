package types

type RawTeamStats struct {
	FGM string `json:"fieldGoalsMade"`
	FGA string `json:"fieldGoalsAttempted"`
	FTM string `json:"freeThrowsMade"`
	FTA string `json:"freeThrowsAttempted"`
	TPM string `json:"threePointsMade"`
	TPA string `json:"threePointsAttempted"`

	OREB string `json:"offensiveRebounds"`
	REB string `json:"totalRebounds"`
	AST string `json:"assists"`
	TO string `json:"turnovers"`
	PF string `json:"personalFouls"`
	STL string `json:"steals"`
	BLK string `json:"blockedShots"`
	PTS string `json:"points"`

	FGpct string `json:"fieldGoalPercentage"`
	TPpct string `json:"threePointPercentage"`
	FTpct string `json:"freeThrowPercentage"`
}

type TeamStats struct {
	FGM uint8
	FGA uint8
	FTM uint8
	FTA uint8
	TPM uint8
	TPA uint8

	OREB uint16
	REB uint16
	AST uint8
	TO uint8
	PF uint8
	STL uint8
	BLK uint8
	PTS uint16

	FGpct float32
	TPpct float32
	FTpct float32
}
