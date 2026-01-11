package types

type RawPeriod struct {
	PeriodDisplay string    `json:"periodDisplay"`
	PlayByPlay    []RawPlay `json:"playbyplayStats"`
}
