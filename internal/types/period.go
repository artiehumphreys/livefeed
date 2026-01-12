package types

type RawPeriod struct {
	PeriodNumber  uint8     `json:"periodNumber"`
	PeriodDisplay string    `json:"periodDisplay"`
	PlayByPlay    []RawPlay `json:"playbyplayStats"`
}

type Period struct {
	PeriodNumber  uint8
	PeriodDisplay string
	PlayByPlay    []Play
}
