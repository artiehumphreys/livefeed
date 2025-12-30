package types

type Team struct {
	TeamID string `json:"teamId"`
	Name string `json:"nameShort"`
	IsHome *string `json:"isHome"`
}
