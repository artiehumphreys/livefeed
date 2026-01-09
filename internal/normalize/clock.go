package normalize

import (
	"math"

	"github.com/artiehumphreys/livefeed/internal/types"
)

func parseClock(raw *types.RawBoxScore) *types.GameClock {
	if raw.Seconds == nil && raw.Minutes == nil {
		return nil
	}

	var minutes, seconds uint8

	if raw.Minutes != nil {
		if m, err := raw.Minutes.Int64(); err == nil {
			minutes = uint8(clampI64(0, math.MaxUint8, m))
		}
	}

	if raw.Seconds != nil {
		if s, err := raw.Seconds.Int64(); err == nil {
			seconds = uint8(clampI64(0, math.MaxUint8, s))
		}
	}

	return &types.GameClock{
		Minutes: minutes,
		Seconds: seconds,
	}
}
