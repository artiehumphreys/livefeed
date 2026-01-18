package normalize

import (
	"math"
	"strings"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func NormalizePlays(raw []types.RawPlay, period uint8) ([]types.Play, error) {
	if len(raw) == 0 {
		return []types.Play{}, nil
	}

	res := make([]types.Play, 0, len(raw))

	for _, p := range raw {
		clock := parseClockFromString(p.Clock)

		teamID, err := p.TeamID.Int64()
		if err != nil {
			teamID = 0
		}

		res = append(res, types.Play{
			Index:            uint16(len(res)),
			TeamID:           uint16(clampI64(0, math.MaxUint16, teamID)),
			Period:           period,
			IsHome:           p.IsHome,
			HomeScore:        p.HomeScore,
			VisitorScore:     p.VisitorScore,
			Clock:            clock,
			FirstName:        strings.TrimSpace(p.FirstName),
			LastName:         strings.TrimSpace(p.LastName),
			EventDescription: strings.TrimSpace(p.EventDescription),
		})
	}

	return res, nil
}
