package normalize

import (
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseTeam(rawBox types.RawTeamBox, rawTeam types.RawTeam) types.Team {
	players := make([]types.PlayerStats, 0, len(rawBox.PlayerStats))
	for _, rp := range rawBox.PlayerStats {
		players = append(players, parsePlayerStats(rp))
	}

	isHome := strings.EqualFold(strings.TrimSpace(rawTeam.IsHome), "true")

	return types.Team{
		TeamID:  stou16(rawTeam.TeamID),
		Name:    strings.TrimSpace(rawTeam.Name),
		IsHome:  isHome,
		Players: players,
		Stats:   parseTeamStats(rawBox.TeamStats),
	}
}
