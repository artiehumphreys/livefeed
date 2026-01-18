package normalize

import (
	"fmt"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func NormalizePeriods(raw []types.RawPeriod) ([]types.Period, error) {
	res := make([]types.Period, 0, len(raw))

	for _, p := range raw {
		playByPlay, err := NormalizePlays(p.PlayByPlay, p.PeriodNumber)
		if err != nil {
			return nil, fmt.Errorf("period %d: %w", p.PeriodNumber, err)
		}

		res = append(res, types.Period{
			PeriodNumber:  p.PeriodNumber,
			PeriodDisplay: p.PeriodDisplay,
			PlayByPlay:    playByPlay,
		})
	}

	return res, nil
}
