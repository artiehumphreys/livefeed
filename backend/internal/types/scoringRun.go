package types

type ActiveScoringRun struct {
	TeamID        uint16
	StartIndex    uint16
	StartPeriod   uint8
	StartClock    *GameClock
	PointsFor     uint16
	PointsAgainst uint16
}

type ScoringRun struct {
	TeamID        uint16
	StartIndex    uint16
	EndIndex      uint16
	PointsFor     uint16
	PointsAgainst uint16

	IsKillShot      bool
	DurationSeconds uint16
	Period          uint8
}
