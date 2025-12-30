package types

type RawTeam struct {
	TeamID string `json:"teamId"`
	Name string `json:"nameShort"`
	IsHome *string `json:"isHome"`
}

type Team struct {
	TeamID uint16
	Name string
	IsHome bool

	Players []PlayerStats
	Stats TeamStats
}
