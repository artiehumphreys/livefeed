package types

type GameStatus string

const (
	StatusLive   GameStatus = "Live"
	StatusFinal  GameStatus = "Final"
	StatusBefore GameStatus = "Before"
)
