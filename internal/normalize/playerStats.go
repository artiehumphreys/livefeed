package normalize

import (
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parsePlayerStats(raw types.RawPlayerStats) types.PlayerStats {
	return types.PlayerStats{
		ID:        stou16(raw.ID),
		Number:    stou8(raw.Number),
		FirstName: strings.TrimSpace(raw.FirstName),
		LastName:  strings.TrimSpace(raw.LastName),
		Position:  strings.TrimSpace(raw.Position),

		MP:  stou8(raw.MP),
		FGM: stou8(raw.FGM),
		FGA: stou8(raw.FGA),
		FTM: stou8(raw.FTM),
		FTA: stou8(raw.FTA),
		TPM: stou8(raw.TPM),
		TPA: stou8(raw.TPA),

		OREB: stou16(raw.OREB),
		REB:  stou16(raw.REB),
		AST:  stou8(raw.AST),
		TO:   stou8(raw.TO),
		PF:   stou8(raw.PF),
		STL:  stou8(raw.STL),
		BLK:  stou8(raw.BLK),
		PTS:  stou16(raw.PTS),
	}
}
